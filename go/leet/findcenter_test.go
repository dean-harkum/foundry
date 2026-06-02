package leet

import (
	// "go/utils"
	"testing"
)

func TestFindCenter(t *testing.T) {
	testCases := []struct {
		edges    [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 3}, {4, 2}}, 2},         // Example use case 1
		{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, 1}, // Example use case 2
		{[][]int{{1, 2}}, -1},                        // Example use case 3
		{[][]int{{1, 2}, {2, 3}}, 2},                 // Example use case 4
		{[][]int{{1, 2}, {3, 4}}, -1},                // Example use case 5
	}

	// Loop over the test cases
	for _, tc := range testCases {
		result := FindCenter(tc.edges)
		if result != tc.expected {
			t.Errorf("ErrorNums(%v) = %v; expected %v", tc.edges, result, tc.expected)
		}
	}
}
