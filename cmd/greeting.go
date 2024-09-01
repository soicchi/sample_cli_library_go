package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(greetingCmd)
}

var greetingCmd = &cobra.Command{
	Use:  "greeting",
	Short: "Prints hello world",
	Long: "This command prints hello world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world")
	},
}
