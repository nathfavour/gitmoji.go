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
			fmt.Print("")
			return
		}
		rand.Seed(time.Now().UnixNano())
		fmt.Print(emojiList[rand.Intn(len(emojiList))].Emoji)
	},
}

var suggestionCmd = &cobra.Command{
	Use:   "suggestion [string]",
	Short: "Suggest an emoji for the given string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		loadEmojis()
		if len(emojiList) == 0 {
			fmt.Print("")
			return
		}
		query := strings.ToLower(args[0])
		for _, e := range emojiList {
			if strings.Contains(strings.ToLower(e.Description), query) {
				fmt.Print(e.Emoji)
				return
			}
			for _, alias := range e.Aliases {
				if strings.Contains(strings.ToLower(alias), query) {
					fmt.Print(e.Emoji)
					return
				}
			}
			for _, tag := range e.Tags {
				if strings.Contains(strings.ToLower(tag), query) {
					fmt.Print(e.Emoji)
					return
				}
			}
		}
		// fallback to random
		rand.Seed(time.Now().UnixNano())
		fmt.Print(emojiList[rand.Intn(len(emojiList))].Emoji)
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
