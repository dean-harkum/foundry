package leet

import (
	"practice/go/base/utils"
	"sort"
)

func FindErrorNums(nums []int) []int {
	sort.Ints(nums)
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	var repeats []int
	for num, count := range counts {
		if count > 1 {
			repeats = append(repeats, num)
			nums = utils.RemoveFromSlice(nums, num)

			for i := 0; i < len(nums); i++ {
				// if the slice index does not match its value, its a repeat
				if nums[i] != i+1 {
					repeats = append(repeats, i+1)
					break
				}
			}
		}
	}
	// if our repeats slice does not have the 'missing number', it was out of bounds per logic above
	if len(repeats) == 1 {
		repeats = append(repeats, len(nums)+1)
	}
	return repeats
}
