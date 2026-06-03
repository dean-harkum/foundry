package cases

import (
	"testing"
)

func TestIsValidParanthesis(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"()", true},     // Example use case 1
		{"()[]{}", true}, // Example use case 2
		{"(]", false},    // Example use case 3
		{"", false},      // Example use case 3
		{"[", false},
		{"]", false},
		{"((", false},
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := IsValidParanthesis(tc.str)
		if result != tc.expected {
			t.Errorf("IsValidParanthesis(%v) = %t; expected %t", tc.str, result, tc.expected)
		}
	}
}
