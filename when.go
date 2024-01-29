package when

import (
	"github.com/the-hidden-layer/when/internal/description"
	"github.com/the-hidden-layer/when/internal/parser"
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
