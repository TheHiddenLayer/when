package when

import (
	"github.com/ubhattac/when/description"
	"github.com/ubhattac/when/parser"
)

func When(time string) (string, error) {
	t, err := parser.Parse(time)
	if err != nil {
		return "", err
	}
	return description.Describe(t), nil
}
