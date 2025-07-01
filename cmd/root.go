package cmd

import (
	"fmt"
	"os"

	"github.com/nathfavour/gitmoji.go/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitmoji",
	Short: "A CLI for using gitmojis in your commit messages",
	Long:  "gitmoji.go is a CLI tool to help you use gitmojis easily in your git workflow.",
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
