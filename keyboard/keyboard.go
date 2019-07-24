package keyboard

import (
	"fmt"
	"log"
	"os"
	"time"
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
	"aqua":   RGBColor{255, 255, 0},
	"blue":   RGBColor{0, 0, 255},
	"pink":   RGBColor{255, 105, 180},
	"purple": RGBColor{128, 0, 128},
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
			go fh.WriteString(color)
			go fh.Close()
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

// InfiniteRainbow generates... an infinite rainbow
func InfiniteRainbow() {
	open := func(files []string) []*os.File {
		kbfiles := make([]*os.File, 0, len(files))
		for _, file := range files {
			p := fmt.Sprintf("/sys/class/leds/system76::kbd_backlight/%v", file)
			fh, err := os.OpenFile(p, os.O_RDWR, 0755)
			if err != nil {
				log.Fatal("error: %v", err)
				continue
			}
			kbfiles = append(kbfiles, fh)
		}
		return kbfiles
	}

	colors := make([]string, 0, 6)
	// generate range of rainbow values
	for i := 0; i <= 255; i++ {
		c := RGBColor{255, i, 0}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{i, 255, 0}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{0, 255, i}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{0, i, 255}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{i, 0, 255}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{255, 0, i}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	kbfiles := open(colorFiles)
	for {
		for _, c := range colors {
			for _, f := range kbfiles {
				f.WriteString(c)
			}
			time.Sleep(time.Nanosecond)
		}
	}
}
