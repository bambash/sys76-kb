package keyboard

import (
	"strconv"
	"time"
)

// BrightnessPulse continuously dials up and down brightness
func BrightnessPulse() {
	for {
		for i := 255; i >= 0; i-- {
			s := strconv.Itoa(i)
			BrightnessFileHandler(s)
			time.Sleep(25 * time.Millisecond)
		}
		for i := 1; i <= 255; i++ {
			s := strconv.Itoa(i)
			BrightnessFileHandler(s)
			time.Sleep(25 * time.Millisecond)
		}
	}
}

// InfiniteRainbow generates... an infinite rainbow
func InfiniteRainbow() {
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

	for {
		for _, c := range colors {
			ColorFileHandler(c)
			time.Sleep(time.Nanosecond)
		}
	}
}
