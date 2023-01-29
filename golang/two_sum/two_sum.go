package two_sum

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func twoPower(num uint64) uint64 {
	var result uint64 = 1

	for power := uint64(1); power <= num; power++ {
		result *= 2
	}

	return result
}

func twoSum(nums []int, target int) []int {

	if len(nums) < 3 {
		return []int{0, 1}[:len(nums)]
	}

	numPossibilities := twoPower(uint64(len(nums))) - 1

	for i := uint64(1); i <= numPossibilities; i++ {
		matches := []int{}
		total := 0
		hits := 0
		for expIndex := uint64(0); (hits < 2) && (expIndex < i); expIndex++ {
			marker := uint64(1) << expIndex
			bitIsSet := marker&i != 0x0

			if bitIsSet {
				hits++
				if hits > 2 {
					break
				}

				matches = append(matches, int(expIndex))
				total += nums[expIndex]
				if (total == target) && (len(matches) == 2) {
					return matches
				}

			}
		}

	}
	return []int{}
}
