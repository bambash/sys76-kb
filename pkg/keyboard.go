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
	"white":  RGBColor{255, 255, 255},
}

var colorFiles = []string{"color", "color_center", "color_left", "color_right", "color_extra"}
var ledClass = []string{"system76_acpi", "system76"}
var sysFSPath =     "/sys/class/leds/%v::kbd_backlight"

type SysPath struct {
    Path  string
    Files [8]string
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

func getSysPath() SysPath {
    ret := SysPath{"",[8]string{}}
	for _, sub := range ledClass {
		d := fmt.Sprintf(sysFSPath, sub)
		if _, err := os.Stat(d); os.IsNotExist(err) != true {
		    ret.Path = d
		    break
	    }
	}
	i := 0
	for _, file := range colorFiles {
		d := fmt.Sprintf("%v/%v", ret.Path, file)
		if _, err := os.Stat(d); os.IsNotExist(err) != true {
		    ret.Files[i] = file
		    i += 1
	    }
	}
	return ret
}

// ColorFileHandler writes a string to colorFiles
func ColorFileHandler(c string) {
    sys := getSysPath()
    if sys.Path == "" {
		log.Fatal("can't get a valid sysfs leds path")
    }
	_, exists := presetColors[c]
	if exists {
		c := presetColors[c]
		color := c.GetColorInHex()
		for _, file := range sys.Files {
		    if file == "" {
		        continue
		    }
			p := fmt.Sprintf("%v/%v", sys.Path, file)
			fh, err := os.OpenFile(p, os.O_RDWR, 0755)
			if err != nil {
				log.Print(err)
				continue
			}
			fh.WriteString(color)
			fh.Close()
		}
	} else {
		for _, file := range sys.Files {
		    if file == "" {
		        continue
		    }
			p := fmt.Sprintf("%v/%v", sys.Path, file)
			fh, err := os.OpenFile(p, os.O_RDWR, 0755)
			if err != nil {
				log.Print(err)
				continue
			}
			fh.WriteString(c)
			fh.Close()
		}
	}
}

// BrightnessFileHandler writes a hex value to brightness, and returns the bytes written
func BrightnessFileHandler(c string) int {
    sys := getSysPath()
    p := fmt.Sprintf("%v/brightness", sys.Path)
	f, err := os.OpenFile(p, os.O_RDWR, 0755)

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
