package parser

import (
	"fmt"
	"time"
)

type ParseTimeFunc func(string) (time.Time, error)

func Parse(input string) (time.Time, error) {
	parseTimeFunc := []ParseTimeFunc{
		parseRFC, parseUnixTime,
	}

	for _, timeParserFunc := range parseTimeFunc {
		t, err := timeParserFunc(input)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time")
}
