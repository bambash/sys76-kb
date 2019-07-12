package cmd

import (
	"fmt"
	"os"

	"github.com/bambash/keyboard/keyboard"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sys76kb",
	Short: "sys76-kb is a keyboard controller for System76 laptops",
	Long: `A simple keyboard contoller built with
				  love by bambash in Go.
				  Complete documentation is available at https://github.com/bambash/sys76-kb`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.ParseFlags(args); err != nil {
			fmt.Printf("Error parsing flags: %v\n", err)
		}

		keyboard.InfiniteRainbow()
	},
}

// Execute is the primary entrypoint for this CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
