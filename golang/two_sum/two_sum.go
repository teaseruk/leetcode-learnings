package two_sum

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func twoSum(nums []int, target int) []int {
	var matches []int
	var copyTarget int
	for startIndex, startValue := range nums {
		copyTarget = target - startValue
		for secondIndex, secondValue := range nums[startIndex+1:] {
			if secondValue == copyTarget {
				return []int{startIndex, startIndex + 1 + secondIndex}
			}
		}
	}

	return matches
}
