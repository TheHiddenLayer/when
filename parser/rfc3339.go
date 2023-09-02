package parser

import (
	"fmt"
	"time"
)

func parseRFCTime(input string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
	}

	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, input)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse RFC time")
}
