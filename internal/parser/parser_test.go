package parser

import "testing"

func TestParser(t *testing.T) {
	tests := []struct {
		inputTime string
		valid     bool
	}{
		{"2006-01-02T15:04:05Z", true},
		{"2006-01-02T15:04:05X", false},
		{"1642782306", true},
		{"1665491898", true},
		{"1639929270", true},
		{"1678113571", true},
		{"1645928405", true},
		{"1654016507", true},
		{"1662091632", true},
		{"1637334002", true},
		{"1669685189", true},
		{"1656784883", true},
		{"some random incorrect format", false},
		{"monday 34 june 8120", false},
	}

	for i, test := range tests {
		_, err := Parse(test.inputTime)
		if err != nil && test.valid {
			t.Errorf("Test case %d failed: expected '%s' to be parsed successfully", i+1, test.inputTime)
		}
		if err == nil && !test.valid {
			t.Errorf("Test case %d failed: expected parsing '%s' to fail", i+1, test.inputTime)
		}
	}
}
