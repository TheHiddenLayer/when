package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ubhattac/when"
)

func main() {
	flag.Usage = usage
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

	fmt.Println(when.When(input))
}

func usage() {
    const userGuide = `

Usage: when <TIME>

        <TIME> can be any standard computer time format.

        Currently supported formats are:
        * RFC3339
        * Unix time

        [Examples]
        $ when 2023-09-01T14:30:00Z
        $ when 1660460492

`
    fmt.Print(userGuide)
}
