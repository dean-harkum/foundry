package cases

import "regexp"

func RomanToInt(s string) int {

	match, _ := regexp.MatchString("^[IVXLCDM]+$", s)

	if !match {
		return 0
	}

	// Lookup table for Roman numerals
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	// Iterate through the Roman numeral from left to right
	for i := 0; i < len(s); i++ {
		currentValue := romanMap[s[i]]

		// If the previous value is less than the current value, subtract twice the previous value
		// This accounts for having already added the previous value in the last iteration
		if prevValue < currentValue {
			total += currentValue - 2*prevValue
		} else {
			total += currentValue
		}

		// Update the previous value
		prevValue = currentValue
	}

	return total
}
