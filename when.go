package when

import (
	"fmt"

	d "github.com/ubhattac/when/description"
	p "github.com/ubhattac/when/parser"
)

func When(time string) (string, error) {
	// Try parsing using RFC3339 parser
	t, err := p.ParseRFC3339(time)
	if err == nil {
		return d.GenerateDescription(t), nil
	}

	// Try parsing using Unix timestamp parser
	t, err = p.ParseUnixTime(time)
	if err == nil {
		return d.GenerateDescription(t), nil
	}

	// If parsing using all parsers fails, return an error
	return "", fmt.Errorf("unable to parse timestamp")
}
