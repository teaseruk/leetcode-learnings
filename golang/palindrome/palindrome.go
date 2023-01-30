package palindrome

func isPalindrome(x int) bool {

	if x >= 0 {
		if x < 10 {
			return true
		}

		order := 0
		topOrder := 1
		for num := x; num >= 10; num /= 10 {
			topOrder *= 10
			order++
		}

		numIters := (order + 1) / 2
		for walker := 0; walker < numIters; walker++ {
			lhs := (x / topOrder) % 10
			rhs := x - 10*(x/10)
			if lhs != rhs {
				return false
			}
			x = x / 10
			topOrder /= 100
		}
		return true
	}

	return false
}
