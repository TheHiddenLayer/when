package when

import (
    "fmt"

    p "github.com/ubhattac/when/parser"
    d "github.com/ubhattac/when/description"
)

func When(input string) (string, error) {
    // Try parsing using RFC3339 parser
    t, err := p.ParseRFC3339(input)
    if err == nil {
        return d.GenerateDescription(t), nil
    }

    // Try parsing using Unix timestamp parser
    t, err = p.ParseUnixTime(input)
    if err == nil {
        return d.GenerateDescription(t), nil
    }

    // If parsing using all parsers fails, return an error
    return "", fmt.Errorf("unable to parse the timestamp")
}

