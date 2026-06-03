package cases

import (
	"sort"
)

func MaximumHappinessSum(happiness []int, k int) int64 {
	// sort the array
	sort.Ints(happiness)
	iterations := 1
	var sum int

	for iterations <= k {
		poppedValue := happiness[len(happiness)-iterations]
		takeHappy := poppedValue - (iterations - 1)
		if takeHappy < 0 {
			takeHappy = 0
		}
		sum = sum + takeHappy
		iterations++
	}

	return int64(sum)
}
