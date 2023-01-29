package two_sum

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)

	var matches []int
	for startIndex := 0; (matches == nil) && (startIndex < lenNums-1); startIndex++ {
		for secondIndex := startIndex + 1; (matches == nil) && (secondIndex < lenNums); secondIndex++ {
			sum := nums[startIndex] + nums[secondIndex]

			if sum == target {
				matches = []int{startIndex, secondIndex}
			}
		}
	}

	return matches
}
