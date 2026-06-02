package leet

import (
	"testing"
)

func TestMaxHappySum(t *testing.T) {
	testCases := []struct {
		happiness []int
		k         int
		expected  int64
	}{
		{[]int{1, 2, 3}, 2, 4},    // Example use case 1
		{[]int{1, 1, 1, 1}, 2, 1}, // Example use case 2
		{[]int{2, 3, 4, 5}, 1, 5}, // Example use case 3
		{[]int{1, 1, 1, 1}, 3, 1}, // Example use case 4
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := MaximumHappinessSum(tc.happiness, tc.k)
		if result != tc.expected {
			t.Errorf("maximumHappinessSum(%v, %d) = %d; expected %d", tc.happiness, tc.k, result, tc.expected)
		}
	}
}
