package two_sum

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello")
}

type Sum struct {
	child *Sum
	value int
}

func twoSum(nums []int, target int) []int {
	if target == 0 {
		return []int{}
	}

	sums := make(map[int]*Sum)

	indices := []int{1, 2}

	numPossibilities := int(math.Pow(2, float64(len(nums))))
	// num possibilities = 2n-1 where n is len(indices)

	currentSlot := 0
	nextSlotChange := 1
	for i := 0; i < numPossibilities; i++ {
		slotValue := nums[currentSlot]
		// create buckets
		sums[currentSlot] = &Sum{child: nil, value: slotValue}
	}

	return indices
}
