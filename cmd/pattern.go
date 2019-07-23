package cmd

import (
	"fmt"

	"github.com/bambash/sys76-kb/keyboard"
	"github.com/spf13/cobra"
)

// Pattern represents keyboard color pattern to run
var Pattern string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&Pattern, "pattern", "p", "rainbow", "the pattern to run (only rainbow for now)")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs a color pattern",
	Long:  `runs a collor pattern that the backlight loops through. (only rainbow currently)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("running infinite rainbow %v\n", Pattern)
		keyboard.InfiniteRainbow()
	},
}
