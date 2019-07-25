package cmd

import (
	"fmt"

	keyboard "github.com/bambash/sys76-kb/pkg"
	"github.com/spf13/cobra"
)

// Pattern represents keyboard color pattern to run
var Pattern string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&Pattern, "pattern", "p", "", "the pattern to run ")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs a backlight pattern",
	Long:  `runs a pattern that the backlight loops through. 'rainbow' or 'pulse'`,
	Run: func(cmd *cobra.Command, args []string) {
		if Pattern != "" {
			if Pattern == "rainbow" {
				fmt.Printf("running pattern %v\n", Pattern)
				keyboard.InfiniteRainbow()
			}
			if Pattern == "pulse" {
				fmt.Printf("running pattern %v\n", Pattern)
				keyboard.BrightnessPulse()
			}
		} else {
			cmd.Help()
		}
	},
}
