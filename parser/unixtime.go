package parser

import (
	"strconv"
	"time"
)

func ParseUnixTime(input string) (time.Time, error) {
    unixTime, err := strconv.ParseInt(input, 10, 64)
    if err != nil {
        return time.Time{}, err
    }
    t := time.Unix(unixTime, 0)
    return t, nil
}
