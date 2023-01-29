package two_sum

import (
	"reflect"
	"testing"
)

func Test_case1(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case2(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case3(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	expected := []int{0, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}
