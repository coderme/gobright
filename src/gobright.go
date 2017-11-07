package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	brightnessFile = "/sys/class/backlight/intel_backlight/brightness"
	maxBrightness = 7000
	minBrightness = 0
)

var (
	value      = flag.Int("value", maxBrightness/2, "Set brightness value, or decr/incr brightness by value based on (-incr/-decr) flags")
	incr       = flag.Bool("incr", false, "Increment brightness percentage by value")
	decr       = flag.Bool("decr", false, "Decrement brightness percentage by value")
	brightness = flag.String("file", brightnessFile, "file where to write new brightness value")
	max        = flag.Int("max", maxBrightness, "Max brightness percentage")
	min        = flag.Int("min", minBrightness, "Min brightness percentage")
	quiet      = flag.Bool("quiet", false, "Be Quiet")
)

func init() {

	flag.Parse()

	if *value == 0 || *incr && *decr {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *max > maxBrightness || *max < 0 || *max <= *min {
		*max = maxBrightness
	}

	if *min < minBrightness || *min >= *max {
		*min = minBrightness
	}

}

func main() {
	bright, err := getBrightness(*brightness)
	if err != nil {
		if !*quiet {
			fmt.Println("[Error]", err)
		}
		os.Exit(1)
	}

	newBright, err2 := setBrightness(bright)

	if !*quiet {

		if err2 != nil {
			fmt.Println("[Error]", err2)
		} else {

			currentVal, _ := getBrightness(*brightness)

			if currentVal != newBright {
				fmt.Println("[INFO] Brightness kept as", currentVal)
			} else {
				fmt.Println("[INFO] Brightness set to", currentVal)
			}
		}

	}

}

func getBrightness(file string) (int, error) {
	oldval, err := ioutil.ReadFile(*brightness)
	if err != nil {
		return 0, err
	}

	bright, err2 := strconv.Atoi(
		strings.TrimSpace(
			string(oldval),
		),
	)
	if err2 != nil {
		return 0, err2
	}

	return bright, nil

}

func setBrightness(oldVal int) (int, error) {
	newVal := oldVal

	if *incr {
		newVal += *value
	}

	if *decr {
		newVal -= *value
	}

	if !*incr && !*decr {
		newVal = *value
	}

	if newVal > *max {
		if !*quiet {
			fmt.Println("[WARNING] Brightness set to 'max'", *max)
		}
		newVal = *max
	}

	if newVal < *min {
		if !*quiet {
			fmt.Println("[WARNING] Brightness set to 'min'", *min)
		}
		newVal = *min
	}
	strVal := strconv.Itoa(newVal)
	return newVal, ioutil.WriteFile(*brightness, []byte(strVal), 0744)

}
