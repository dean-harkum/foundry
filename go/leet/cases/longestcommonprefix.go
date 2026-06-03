package cases

import (
	"github.com/dean-harkum/foundry/go/leet/utils"
)

func LongestCommonPrefix(strs []string) string {
	CommonLetters := ""
	ShortestString := utils.FindShortestString(strs)

	for i := 0; i < len(ShortestString); i++ {
		letters := make(map[string]int)
		for j := 0; j < len(strs); j++ {
			letters[string(strs[j][i])]++
		}
		if len(letters) == 1 {
			for key := range letters {
				CommonLetters = CommonLetters + key
			}
		} else {
			return CommonLetters
		}
	}
	return CommonLetters
}
