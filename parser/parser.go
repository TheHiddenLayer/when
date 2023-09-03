package parser

import (
	"fmt"
	"time"
)

type TimeParser func(string) (time.Time, error)

func Parse(input string) (time.Time, error) {
	timeParsers := []TimeParser{parseRFCTime, parseUnixTime}

	for _, parseTime := range timeParsers {
		parsedTime, err := parseTime(input)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time")
}
