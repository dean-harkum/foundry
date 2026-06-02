package leet

import (
	"go/utils"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}}, // Example use case 1
		{[]int{3, 2, 4}, 6, []int{1, 2}},      // Example use case 2
		{[]int{3, 3}, 6, []int{0, 1}},         // Example use case 3
		{[]int{3}, 6, []int{}},                // Example use case 4
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := TwoSum(tc.nums, tc.target)
		if !utils.SlicesEqual(result, tc.expected) {
			t.Errorf("twoSums(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
		}
	}
}
