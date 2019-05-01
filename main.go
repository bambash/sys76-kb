package main

import (
	"fmt"
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
func GetColorInHex(c RGBColor) string {
	hex := getHex(c.Red) + getHex(c.Green) + getHex(c.Blue)
	return hex
}

func main() {

	// first loop
	for i := 0; i <= 255; i++ {
		c := RGBColor{255, i, 0}
		fmt.Println(GetColorInHex(c))
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{i, 255, 0}
		fmt.Println(GetColorInHex(c))
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{0, 255, i}
		fmt.Println(GetColorInHex(c))
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{0, i, 255}
		fmt.Println(GetColorInHex(c))
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{i, 0, 255}
		fmt.Println(GetColorInHex(c))
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{255, 0, i}
		fmt.Println(GetColorInHex(c))
	}
}
