package keyboard

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// BrightnessPulse continuously dials up and down brightness
func BrightnessPulse() {
	for {
		for i := 255; i >= 0; i-- {
			s := strconv.Itoa(i)
			BrightnessFileHandler(s)
			time.Sleep(10 * time.Millisecond)
		}
		for i := 0; i <= 255; i++ {
			s := strconv.Itoa(i)
			BrightnessFileHandler(s)
			time.Sleep(10 * time.Millisecond)
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
		colors = append(colors, c.GetColorInHex())
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{i, 255, 0}
		colors = append(colors, c.GetColorInHex())
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{0, 255, i}
		colors = append(colors, c.GetColorInHex())
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{0, i, 255}
		colors = append(colors, c.GetColorInHex())
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{i, 0, 255}
		colors = append(colors, c.GetColorInHex())
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{255, 0, i}
		colors = append(colors, c.GetColorInHex())
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
