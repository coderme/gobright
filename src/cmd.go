package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	name    = "GoBright"
	version = "2.0"
	license = `
Copyright (c) %d CoderMe.com
 Permission to use, copy, modify, and distribute this software for any
 purpose with or without fee is hereby granted, provided that the above
 copyright notice and this permission notice appear in all copies.
 
 THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
`
 	brightnessFile = "/sys/class/backlight/intel_backlight/brightness"
	maxBrightness = 7000
	minBrightness = 0
 
)

var (
	value      = flag.Int("v", maxBrightness/2, "")
	incr       = flag.Bool("i", false, "")
	decr       = flag.Bool("d", false, "")
	brightness = flag.String("f", brightnessFile, "")
	max        = flag.Int("x", maxBrightness, "")
	min        = flag.Int("m", minBrightness, "")
	quiet      = flag.Bool("q", false, "")
	
	help        = flag.Bool("h", false, "")
	showhelp    = flag.Bool("help", false, "")
	showLicense = flag.Bool("l", false, "")
	showVersion = flag.Bool("version", false, "")
)


 
func printVersion() {
	fmt.Println(getVersion())
	os.Exit(0)

}

func getVersion() string {
	return name + " v" + version

}

func printLicense() {
	l := "\n" + getVersion() + "\n"
	l += license + "\n"
	fmt.Printf(l, time.Now().Year())
	os.Exit(0)

}

func usage() {
	fmt.Printf(`%s

Usage: %s [-l | -version -h | --help] [(-d | -i)] [-f file] [-x max] [-m min] [-v value]

FLAGS:
 -version
    Show version and exit.
 -l 
    Show License and exit.
 -h | --help
    Show help and exit.
 -d
    Decrement brightness percentage by value
 -i
    Increment brightness percentage by value
 -q 
    Be Quiet


OPTIONS:
 -v value
    Set brightness value, or decr/incr brightness by value based on (-d/-i) flags, (default: %d)
 -f file
    File location where to write new brightness value, (default: %s)
 -x max
   Max brightness percentage, (default: %d)
 -m min
   Min brightness percentage, (default: %d)

`, getVersion(), os.Args[0], maxBrightness/2,
		brightnessFile, maxBrightness, minBrightness)

	os.Exit(1)

}

