package parser

import (
	"strconv"
	"time"
)

func parseUnixTime(input string) (time.Time, error) {
	unixTime, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	parsedTime := time.Unix(unixTime, 0)
	return parsedTime, nil
}
