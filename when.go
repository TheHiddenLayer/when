package when

import (
	"github.com/ubhattac/when/description"
	"github.com/ubhattac/when/parser"
)

func When(time string) (string, error) {
	parsedTime, err := parser.Parse(time)
	if err != nil {
		return "", err
	}
	return description.Describe(parsedTime), nil
}

func WhenVerbosely(time string) (string, error) {
	parsedTime, err := parser.Parse(time)
	if err != nil {
		return "", err
	}
	return description.DescribeVerbosely(parsedTime), nil
}
