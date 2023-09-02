package parser

import (
    "time"
)

func ParseRFC3339(input string) (time.Time, error) {
    t, err := time.Parse(time.RFC3339, input)
    if err != nil {
        return time.Time{}, err
    }
    return t, nil
}

