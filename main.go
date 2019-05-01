package main

import (
	"fmt"
	"log"
	"os"
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

func FileHandler(c string) bool {
	f, err := os.OpenFile("/sys/class/leds/system76::kbd_backlight/color_left", os.O_RDWR, 0755)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	l, err := f.WriteString(c)
	if err != nil {
		log.Fatal(err)
		f.Close()
		return false
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func main() {

	// infinite rainbow
	for {
		for i := 0; i <= 255; i++ {
			c := RGBColor{255, i, 0}
			fmt.Println(GetColorInHex(c))
			hex := GetColorInHex(c)
			FileHandler(hex)
		}

		for i := 255; i >= 0; i-- {
			c := RGBColor{i, 255, 0}
			fmt.Println(GetColorInHex(c))
			hex := GetColorInHex(c)
			FileHandler(hex)
		}

		for i := 0; i <= 255; i++ {
			c := RGBColor{0, 255, i}
			fmt.Println(GetColorInHex(c))
			hex := GetColorInHex(c)
			FileHandler(hex)
		}

		for i := 255; i >= 0; i-- {
			c := RGBColor{0, i, 255}
			fmt.Println(GetColorInHex(c))
			hex := GetColorInHex(c)
			FileHandler(hex)
		}

		for i := 0; i <= 255; i++ {
			c := RGBColor{i, 0, 255}
			fmt.Println(GetColorInHex(c))
			hex := GetColorInHex(c)
			FileHandler(hex)
		}

		for i := 255; i >= 0; i-- {
			c := RGBColor{255, 0, i}
			fmt.Println(GetColorInHex(c))
			hex := GetColorInHex(c)
			FileHandler(hex)
		}
	}
}
