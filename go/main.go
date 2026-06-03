package main

import (
	"fmt"
	"github.com/dean-harkum/foundry/go/leet"
)

func main() {
	// result := leet.IsPalindrome(121)
	// result := leet.RomanToInt("III")
	// result := leet.FindErrorNums([]int{1,5,3,2,2,7,6,4,8,9})
	// result := leet.LongestCommonPrefix([]string{"flower","flow","flight"})
	// result := leet.IsValidParanthesis("[")
	result := leet.FindCenter([][]int{{1, 2}, {2, 3}, {4, 2}})

	fmt.Println(result)
	fmt.Println()
}
