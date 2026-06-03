package main

import (
	"fmt"
	"github.com/dean-harkum/foundry/go/leet/cases"
)

func main() {
	// result := cases.IsPalindrome(121)
	// result := cases.RomanToInt("III")
	// result := cases.FindErrorNums([]int{1,5,3,2,2,7,6,4,8,9})
	// result := cases.LongestCommonPrefix([]string{"flower","flow","flight"})
	// result := cases.IsValidParanthesis("[")
	result := cases.FindCenter([][]int{{1, 2}, {2, 3}, {4, 2}})

	fmt.Println(result)
	fmt.Println()
}
