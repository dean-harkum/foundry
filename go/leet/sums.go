package leet

func TwoSum(nums []int, target int) []int {

	var twosum []int
	for index, element := range nums {
		num1 := element
		for index2, element2 := range nums {
			if index2 != index {
				if num1+element2 == target {
					twosum = append(twosum, index, index2)
					return twosum
				}
			}
		}
	}
	return twosum
}
