package cmd

import (
	"fmt"

	"github.com/nathfavour/gitmoji.go/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available emojis",
	Run: func(cmd *cobra.Command, args []string) {
		emojis, err := config.LoadEmojis()
		if err != nil {
			fmt.Println("Error loading emojis:", err)
			return
		}
		for _, e := range emojis {
			fmt.Printf("%s  %s  :%s:\n", e.Emoji, e.Description, e.Aliases[0])
		}
	},
}
