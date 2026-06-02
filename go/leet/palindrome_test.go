package leet

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		number   int
		expected bool
	}{
		{121, true},   // Example use case 1
		{-121, false}, // Example use case 2
		{10, false},   // Example use case 3
		{12321, true}, // Example use case 4
		{123, false},  // Example use case 5
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := IsPalindrome(tc.number)
		if result != tc.expected {
			t.Errorf("isPalindrome(%d) = %t; expected %t", tc.number, result, tc.expected)
		}
	}
}
