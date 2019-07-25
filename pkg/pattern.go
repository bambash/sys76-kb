package keyboard

import (
	"fmt"
	"log"
	"os"
	"time"
)

// BrightnessPulse continuously dials up and down brightness
func BrightnessPulse() {
	f, err := os.OpenFile("/sys/class/leds/system76::kbd_backlight/brightness", os.O_RDWR, 0755)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for {
		for i := 255; i >= 0; i-- {
			s := string(i)
			f.WriteString(s)
			time.Sleep(100 * time.Millisecond)
		}
		for i := 0; i <= 255; i++ {
			s := string(i)
			f.WriteString(s)
			time.Sleep(100 * time.Millisecond)
		}
	}
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
