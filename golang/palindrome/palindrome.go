package palindrome

import "math"

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

		numIters := int(math.Max(1, float64(order)/2.0))

		for walker := 0; walker < numIters; walker++ {
			lhs := x / topOrder
			rhs := x - 10*(x/10)
			if lhs != rhs {
				return false
			}
			x = (x - lhs*topOrder) / 10
			topOrder /= 100
		}
		return true
	}

	return false
}
