package description

import (
	"fmt"
	"testing"
	"time"
)

type descriptionTest struct {
	inputTime           time.Time
	expectedDescription string
}

func TestDescribe(t *testing.T) {
	baseTime := time.Now()
	tests := []descriptionTest{
		{baseTime.Add(-5 * second), "now"},
		{baseTime.Add(20 * second), "now"},
		{baseTime.Add(-21 * second), "21s"},
		{baseTime.Add(-59 * second), "59s"},
		{baseTime.Add(1*minute + 10*second), "in 1m"},
		{baseTime.Add(-(1*minute + 21*second)), "1m"},
		{baseTime.Add(-2 * hour), "2h"},
		{baseTime.Add(2 * day), "in 2d"},
		{baseTime.Add(-3 * week), "3w"},
		{baseTime.Add(4 * month), fmt.Sprintf("on %s", baseTime.Add(4*month).Format("01/02/2006"))},
		{baseTime.Add(-4 * year), baseTime.Add(-4 * year).Format("01/02/2006")},
	}

	for i, test := range tests {
		description := Describe(test.inputTime)
		if description != test.expectedDescription {
			t.Errorf("Test case %d failed: expected '%s' but got '%s'", i+1, test.expectedDescription, description)
		}
	}
}

func TestDescribeVerbosely(t *testing.T) {
	baseTime := time.Now()
	tests := []descriptionTest{
		{baseTime.Add(-5 * second), "just a few seconds ago"},
		{baseTime.Add(20 * second), "in less than 1 minute"},
		{baseTime.Add(-21 * second), "less than 1 minute ago"},
		{baseTime.Add(-59 * second), "about 1 minute ago"},
		{baseTime.Add(1*minute + 20*second), "in 1 minute"},
		{baseTime.Add(-(1*minute + 21*second)), "1 minute ago"},
		{baseTime.Add(-2 * hour), "2 hours ago"},
		{baseTime.Add(2*day + hour), "in almost 2 days"},
		{baseTime.Add(-3 * week), "around 3 weeks ago"},
		{baseTime.Add(4*30*24*time.Hour + time.Minute), "in roughly 4 months"},
		{baseTime.Add(-4 * year), "over 4 years ago"},
	}

	for i, test := range tests {
		description := DescribeVerbosely(test.inputTime)
		if description != test.expectedDescription {
			t.Errorf("Test case %d failed: expected '%s' but got '%s'", i+1, test.expectedDescription, description)
		}
	}
}
