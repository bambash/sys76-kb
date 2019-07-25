package cmd

import (
	keyboard "github.com/bambash/sys76-kb/pkg"
	"github.com/spf13/cobra"
)

// Pattern represents keyboard color pattern to run
var Color string
var Brightness string

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&Color, "color", "c", "", "hex or string value of a color (red, blue, green, purple, pink, orange, yellow")
	setCmd.Flags().StringVarP(&Brightness, "brightness", "b", "", "sets the backlight brightness (0 - 255")
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the color and brightness of the keyboard",
	Long:  `Sets the color and brightness of the keyboard`,
	Run: func(cmd *cobra.Command, args []string) {
		if Color != "" {
			keyboard.ColorFileHandler(Color)
		}
		if Brightness != "" {
			keyboard.BrightnessFileHandler(Brightness)
		}
	},
}
