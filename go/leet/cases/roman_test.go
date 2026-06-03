package cases

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		number   string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"XMCMX0LV", 0},
	}

	for _, tc := range testCases {
		result := RomanToInt(tc.number)
		if result != tc.expected {
			t.Errorf("romanToInit(%s) = %d; expected %d", tc.number, result, tc.expected)
		}
	}
}
