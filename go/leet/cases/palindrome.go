package cases

func IsPalindrome(x int) bool {
	// return false before doing anything if negative
	if x < 0 {
		return false
	}

	// save the original number, and init reversed
	original := x
	reversed := 0

	for x != 0 {
		// grab the last number, and append to our reversed var
		reversed = reversed*10 + x%10
		// divide by 10 to remove the last digit
		x /= 10
	}

	return original == reversed
}
