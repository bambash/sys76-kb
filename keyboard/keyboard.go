package keyboard

import (
	"fmt"
	"log"
	"os"
	"time"
)

type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

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

// WriteColorFile writes a hex value to "f" input, and returns the bytes written
func WriteColorFile(c string, f string) int {
	path := fmt.Sprintf("/sys/class/leds/system76::kbd_backlight/%v", f)
	kbfile, err := os.OpenFile(path, os.O_RDWR, 0755)

	if err != nil {
		log.Fatal(err)
	}

	l, err := kbfile.WriteString(c)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	kbfile.Close()
	return l
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
	files := []string{"color_center", "color_left", "color_right", "color_extra"}

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

	kbfiles := open(files)
	for {
		for _, c := range colors {
			for _, f := range kbfiles {
				f.WriteString(c)
			}
			time.Sleep(time.Nanosecond)
		}
	}
}
