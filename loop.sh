#!/usr/bin/env bash
while true; do
    for i in $(go run main.go); do 
        echo $i > /sys/class/leds/system76\:\:kbd_backlight/color_left; 
    done;
done;