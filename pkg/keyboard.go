package keyboard

import (
	"fmt"
	"log"
	"os"
)

// RGBColor represents Red Green and Blue values of a color
type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

var presetColors = map[string]RGBColor{
	"red":    RGBColor{255, 0, 0},
	"orange": RGBColor{255, 128, 0},
	"yellow": RGBColor{255, 255, 0},
	"green":  RGBColor{0, 255, 0},
	"aqua":   RGBColor{25, 255, 223},
	"blue":   RGBColor{0, 0, 255},
	"pink":   RGBColor{255, 105, 180},
	"purple": RGBColor{128, 0, 128},
	"white":  RGBColor{0, 0, 0},
}

var colorFiles = []string{"color_center", "color_left", "color_right", "color_extra"}

// GetHex Converts a decimal number to hex representations
func getHex(num int) string {
	hex := fmt.Sprintf("%X", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

// GetColorInHex returns a color in HEX format
func (c RGBColor) GetColorInHex() string {
	hex := getHex(c.Red) + getHex(c.Green) + getHex(c.Blue)
	return hex
}

// ColorFileHandler writes a hex value to "f" input, and returns the bytes written
func ColorFileHandler(c string) {
	_, exists := presetColors[c]
	if exists {
		c := presetColors[c]
		color := c.GetColorInHex()
		for _, file := range colorFiles {
			p := fmt.Sprintf("/sys/class/leds/system76::kbd_backlight/%v", file)
			fh, err := os.OpenFile(p, os.O_RDWR, 0755)
			if err != nil {
				log.Fatal("error: %v", err)
				os.Exit(1)
			}
			fh.WriteString(color)
			fh.Close()
		}
	} else {
		os.Exit(1)
	}
}

// BrightnessFileHandler writes a hex value to brightness, and returns the bytes written
func BrightnessFileHandler(c string) int {
	f, err := os.OpenFile("/sys/class/leds/system76::kbd_backlight/brightness", os.O_RDWR, 0755)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	l, err := f.WriteString(c)
	if err != nil {
		log.Fatal(err)
		f.Close()
		return 0
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return l
}
