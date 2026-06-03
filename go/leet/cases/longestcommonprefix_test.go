package cases

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"}, // Example use case 1
		{[]string{"dog", "racecar", "car"}, ""},      // Example use case 2
		{[]string{"cir", "car"}, "c"},                // Example use case 3
		{[]string{}, ""},                             // Example use case 3
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := LongestCommonPrefix(tc.strs)
		if result != tc.expected {
			t.Errorf("LongestCommonPrefix(%v) = %s; expected %s", tc.strs, result, tc.expected)
		}
	}
}
