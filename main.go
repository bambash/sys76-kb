package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

// HSVColor HSV Color Type
type HSVColor struct {
	Hue        float64
	Saturation float64
	Value      float64
}

// GetHex Converts a decimal number to hex representations
func getHex(num int) string {
	hex := fmt.Sprintf("%X", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func rgbToHSV(color RGBColor) HSVColor {
	max := math.Max(float64(color.Red), float64(color.Green))
	max = math.Max(max, float64(color.Blue))
	min := math.Min(float64(color.Red), float64(color.Green))
	min = math.Min(min, float64(color.Blue))
	delta := max - min

	red := float64(color.Red)
	green := float64(color.Green)
	blue := float64(color.Blue)

	var hue, saturation, value float64

	if delta == 0 {
		hue = 0
	} else if red == max {
		hue = float64(int(((green - blue) / delta)) % 6)
	} else if green == max {
		hue = (blue-red)/delta + 2
	} else if blue == max {
		hue = (red-green)/delta + 4
	}

	hue = math.Round(hue * 60)
	if hue < 0 {
		hue += 360
	}

	if max == 0 {
		saturation = 0
	} else {
		saturation = (delta / max) * 100
	}

	value = max / 255 * 100

	return HSVColor{hue, saturation, value}
}

// GetRandomColorInRgb Returns a random RGBColor
func GetRandomColorInRgb() RGBColor {
	rand.Seed(time.Now().UnixNano())
	Red := rand.Intn(255)
	Green := rand.Intn(255)
	blue := rand.Intn(255)
	c := RGBColor{Red, Green, blue}
	return c
}

// GetRandomColorInHex returns a random color in HEX format
func GetRandomColorInHex() string {
	color := GetRandomColorInRgb()
	hex := getHex(color.Red) + getHex(color.Green) + getHex(color.Blue)
	return hex
}

// GetRandomColorInHSV returns a random color in HSV format
func GetRandomColorInHSV() HSVColor {
	color := GetRandomColorInRgb()
	return rgbToHSV(color)
}

// GetColorInHex returns a color in HEX format
func GetColorInHex(c RGBColor) string {
	hex := getHex(c.Red) + getHex(c.Green) + getHex(c.Blue)
	return hex
}

// RGBdiff returns the difference bettween two RGBColor inputs
func RGBdiff(c1 RGBColor, c2 RGBColor) RGBColor {

	r := (c1.Red + c2.Red) / 2
	g := (c1.Green + c2.Green) / 2
	b := (c1.Blue + c2.Blue) / 2

	avg := RGBColor{r, g, b}
	return avg
}

func RGBGradient(c1 RGBColor, c2 RGBColor, steps float64) []RGBColor {

	var gradient []RGBColor

	gradient = append(gradient, c1)

	rInc := int(math.Round(float64(c2.Red-c1.Red) / (steps + 1)))
	gInc := int(math.Round(float64(c2.Green-c1.Green) / (steps + 1)))
	bInc := int(math.Round(float64(c2.Blue+c1.Blue) / (steps + 1)))

	for i := 0; i < int(steps); i++ {
		c1.Red += rInc
		c1.Green += gInc
		c1.Blue += bInc

		if c1.Red >= 255 {
			continue
		}
		if c1.Green >= 255 {
			continue
		}
		if c1.Blue >= 255 {
			continue
		}
		newRGB := RGBColor{c1.Red, c2.Green, c1.Blue}
		gradient = append(gradient, newRGB)
	}

	gradient = append(gradient, c2)

	return gradient
}

func main() {
	violet := RGBColor{148, 0, 211}
	//indigo := RGBColor{75, 0, 130}
	blue := RGBColor{0, 0, 255}
	//green := RGBColor{0, 255, 0}
	yellow := RGBColor{255, 255, 0}
	orange := RGBColor{255, 127, 0}
	red := RGBColor{255, 0, 0}

	vb := RGBGradient(violet, blue, 50)
	by := RGBGradient(yellow, blue, 50)
	yo := RGBGradient(yellow, orange, 2)
	or := RGBGradient(orange, red, 2)
	rv := RGBGradient(red, violet, 50)

	for _, v := range vb {
		fmt.Println(GetColorInHex(v))
	}

	for _, v := range by {
		fmt.Println(GetColorInHex(v))
	}

	for _, v := range yo {
		fmt.Println(GetColorInHex(v))
	}

	for _, v := range or {
		fmt.Println(GetColorInHex(v))
	}

	for _, v := range rv {
		fmt.Println(GetColorInHex(v))
	}
	// for i := 0; i <= 255; i++ {
	// 	color := RGBColor{0, i, 0)
	// 	fmt.Println(color)
	// }

}
