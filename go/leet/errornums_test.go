package leet

import (
	"go/utils"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},                    // Example use case 1
		{[]int{1, 1}, []int{1, 2}},                          // Example use case 2
		{[]int{2, 2}, []int{2, 1}},                          // Example use case 3
		{[]int{3, 2, 2}, []int{2, 1}},                       // Example use case 4
		{[]int{1, 3, 3}, []int{3, 2}},                       // Example use case 5
		{[]int{3, 2, 3, 4, 6, 5}, []int{3, 1}},              // Example use case 6
		{[]int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}, []int{2, 10}}, // Example use case 7
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := FindErrorNums(tc.nums)
		if !utils.SlicesEqual(result, tc.expected) {
			t.Errorf("ErrorNums(%v) = %v; expected %v", tc.nums, result, tc.expected)
		}
	}
}
