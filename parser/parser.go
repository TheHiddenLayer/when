package parser

import (
	"fmt"
	"time"
)

type TimeParserFunc func(string) (time.Time, error)

func Parse(input string) (time.Time, error) {

	timeParserFuncs := []TimeParserFunc{
		ParseRFC3339, ParseUnixTime,
	}

	for _, timeParserFunc := range timeParserFuncs {
		t, err := timeParserFunc(input)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time")
}
