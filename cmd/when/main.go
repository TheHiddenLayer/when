package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ubhattac/when"
)

func main() {
	flag.Usage = usage
	verbose := flag.Bool("v", false, "Enable verbose mode to display wordy time descriptions")
	flag.Parse()

	input := flag.Arg(0)

	if input == "help" {
		usage()
		os.Exit(0)
	}

	if input == "" {
		usage()
		os.Exit(1)
	}

	var readableTime string
	var err error
	if *verbose {
		readableTime, err = when.WhenVerbosely(input)
	} else {
		readableTime, err = when.When(input)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(readableTime)
}

func usage() {
	const userGuide = `
Usage: when [-v] <timestamp>

        <timestamp> can be any standard computer time format.

        Currently supported formats are:
        * RFC formats
        * Unix timestamp

        [Examples]
    		when 2023-09-01T14:30:00Z
    		when -v 1660460492

`
	fmt.Print(userGuide)
}
