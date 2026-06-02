package leet

func FindCenter(edges [][]int) int {

	// if length of arrays is less than 2, bail
	if len(edges) < 2 {
		return -1
	}

	slice1 := edges[0]
	slice2 := edges[1]

	// Create a map to store the elements of the first slice
	elementsMap := make(map[int]bool)

	// Populate the map with elements from the first slice
	for _, num := range slice1 {
		elementsMap[num] = true
	}

	// Check elements from the second slice against the map
	for _, num := range slice2 {
		if elementsMap[num] {
			return num
		}
	}

	return -1
}
