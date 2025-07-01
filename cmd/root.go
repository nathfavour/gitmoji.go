package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nathfavour/gitmoji.go/internal/config"
	"github.com/spf13/cobra"
)

type Emoji struct {
	Emoji       string   `json:"emoji"`
	Description string   `json:"description"`
	Aliases     []string `json:"aliases"`
	Tags        []string `json:"tags"`
	// ...other fields...
}

var emojiList []Emoji

func loadEmojis() {
	// Load emoji list from sources/default/emoji.json
	home, _ := os.UserHomeDir()
	emojiPath := filepath.Join(home, ".gitmojigo", "sources", "default", "emoji.json")
	data, err := os.ReadFile(emojiPath)
	if err != nil {
		emojiList = []Emoji{}
		return
	}
	_ = json.Unmarshal(data, &emojiList)
}

var rootCmd = &cobra.Command{
	Use:   "gitmoji",
	Short: "A CLI for using gitmojis in your commit messages",
	Long:  "gitmoji.go is a CLI tool to help you use gitmojis easily in your git workflow.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gitmoji.go! Use -h for help.")
	},
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Output a random emoji",
	Run: func(cmd *cobra.Command, args []string) {
		loadEmojis()
		if len(emojiList) == 0 {
			return
		}
		rand.Seed(time.Now().UnixNano())
		os.Stdout.WriteString(emojiList[rand.Intn(len(emojiList))].Emoji + "\n")
	},
}

func normalizeWord(word string) string {
	// Simple stemming - remove common suffixes
	word = strings.TrimSuffix(word, "ing")
	word = strings.TrimSuffix(word, "ed")
	word = strings.TrimSuffix(word, "er")
	word = strings.TrimSuffix(word, "ly")
	word = strings.TrimSuffix(word, "s")
	return word
}

func fuzzyMatch(s1, s2 string) bool {
	// Simple character overlap check (at least 70% common chars)
	if len(s1) < 2 || len(s2) < 2 {
		return false
	}
	common := 0
	for _, c := range s1 {
		if strings.ContainsRune(s2, c) {
			common++
		}
	}
	return float64(common)/float64(len(s1)) >= 0.7
}

var suggestionCmd = &cobra.Command{
	Use:   "suggestion [string]",
	Short: "Suggest an emoji for the given string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		loadEmojis()
		if len(emojiList) == 0 {
			return
		}
		query := strings.ToLower(strings.TrimSpace(args[0]))
		queryNorm := normalizeWord(query)

		// 1. Exact match on alias, tag, description
		for _, e := range emojiList {
			for _, alias := range e.Aliases {
				if strings.ToLower(alias) == query {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			for _, tag := range e.Tags {
				if strings.ToLower(tag) == query {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			if strings.ToLower(e.Description) == query {
				os.Stdout.WriteString(e.Emoji + "\n")
				return
			}
		}

		// 2. Normalized/stemmed exact match
		for _, e := range emojiList {
			for _, alias := range e.Aliases {
				if normalizeWord(strings.ToLower(alias)) == queryNorm {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			for _, tag := range e.Tags {
				if normalizeWord(strings.ToLower(tag)) == queryNorm {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			if normalizeWord(strings.ToLower(e.Description)) == queryNorm {
				os.Stdout.WriteString(e.Emoji + "\n")
				return
			}
		}

		// 3. Prefix match on alias, tag, description
		for _, e := range emojiList {
			for _, alias := range e.Aliases {
				if strings.HasPrefix(strings.ToLower(alias), query) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			for _, tag := range e.Tags {
				if strings.HasPrefix(strings.ToLower(tag), query) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			if strings.HasPrefix(strings.ToLower(e.Description), query) {
				os.Stdout.WriteString(e.Emoji + "\n")
				return
			}
		}

		// 4. Word boundary match (query as whole word)
		queryWord := " " + query + " "
		for _, e := range emojiList {
			if strings.Contains(" "+strings.ToLower(e.Description)+" ", queryWord) {
				os.Stdout.WriteString(e.Emoji + "\n")
				return
			}
			for _, alias := range e.Aliases {
				if strings.Contains(" "+strings.ToLower(alias)+" ", queryWord) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			for _, tag := range e.Tags {
				if strings.Contains(" "+strings.ToLower(tag)+" ", queryWord) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
		}

		// 5. Fuzzy character match
		for _, e := range emojiList {
			for _, alias := range e.Aliases {
				if fuzzyMatch(query, strings.ToLower(alias)) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			for _, tag := range e.Tags {
				if fuzzyMatch(query, strings.ToLower(tag)) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			if fuzzyMatch(query, strings.ToLower(e.Description)) {
				os.Stdout.WriteString(e.Emoji + "\n")
				return
			}
		}

		// 6. Substring match on alias, tag, description (as before)
		for _, e := range emojiList {
			if strings.Contains(strings.ToLower(e.Description), query) {
				os.Stdout.WriteString(e.Emoji + "\n")
				return
			}
			for _, alias := range e.Aliases {
				if strings.Contains(strings.ToLower(alias), query) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
			for _, tag := range e.Tags {
				if strings.Contains(strings.ToLower(tag), query) {
					os.Stdout.WriteString(e.Emoji + "\n")
					return
				}
			}
		}

		// fallback to random
		rand.Seed(time.Now().UnixNano())
		os.Stdout.WriteString(emojiList[rand.Intn(len(emojiList))].Emoji + "\n")
	},
}

func Execute() {
	// Ensure config and sources exist before any command
	if err := config.EnsureConfig(); err != nil {
		fmt.Println("Error initializing config:", err)
		os.Exit(1)
	}
	if err := config.EnsureDefaultSource(); err != nil {
		fmt.Println("Error ensuring default source:", err)
		os.Exit(1)
	}
	if err := config.EnsureDefaultEmojiJSON(); err != nil {
		fmt.Println("Error downloading emoji list:", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(randomCmd)
	rootCmd.AddCommand(suggestionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
func init() {
	rootCmd.AddCommand(listCmd)
}
