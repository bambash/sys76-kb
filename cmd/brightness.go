package cmd

import (
	"fmt"

	"github.com/bambash/sys76-kb/keyboard"

	"github.com/spf13/cobra"
)

// Lumens is a value from 0 - 255 that represents keyboard brightness
var Lumens string

func init() {
	rootCmd.AddCommand(brightnessCmd)
	brightnessCmd.Flags().StringVarP(&Lumens, "lumens", "l", "", "value between 0 and 255")
	brightnessCmd.MarkFlagRequired("lumens")
}

var brightnessCmd = &cobra.Command{
	Use:   "brightness",
	Short: "sets the brightness of the keyboard",
	Long:  `sets the brightness of the keyboard backlight.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("setting brightness to %v\n", Lumens)
		keyboard.BrightnessFileHandler(Lumens)
	},
}
