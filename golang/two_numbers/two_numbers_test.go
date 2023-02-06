package two_numbers

import (
	"reflect"
	"testing"
)

func buildChain(numbers []int) *ListNode {

	var firstChain *ListNode = nil
	var prevNode *ListNode = nil
	for i, num := range numbers {
		if i == 0 {
			firstChain = &ListNode{Val: num, Next: nil}
			prevNode = firstChain
		} else {
			prevNode.Next = &ListNode{Val: num, Next: nil}
			prevNode = prevNode.Next
		}
	}
	return firstChain
}

func extractChain(listNode *ListNode) []int {

	var results []int
	for listNode != nil {
		results = append(results, listNode.Val)
		listNode = listNode.Next
	}
	return results
}

func Test_testBuildChain(t *testing.T) {
	nodes := buildChain([]int{2, 4, 3})

	for nodes != nil {
		println(nodes.Val)
		nodes = nodes.Next
	}
}

func Test_testcase1(t *testing.T) {
	result := addTwoNumbers(buildChain([]int{2, 4, 3}), buildChain([]int{5, 6, 4}))
	expected := []int{7, 0, 8}

	if !reflect.DeepEqual(extractChain(result), expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testcase3(t *testing.T) {
	result := addTwoNumbers(buildChain([]int{0}), buildChain([]int{0}))
	expected := []int{0}

	if !reflect.DeepEqual(extractChain(result), expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testcase4(t *testing.T) {
	result := addTwoNumbers(buildChain([]int{2, 4, 9}), buildChain([]int{5, 6, 4, 9}))
	expected := []int{7, 0, 4, 0, 1}

	if !reflect.DeepEqual(extractChain(result), expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testcase5(t *testing.T) {
	result := addTwoNumbers(buildChain([]int{9, 9, 9, 9, 9, 9, 9}), buildChain([]int{9, 9, 9, 9}))
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	if !reflect.DeepEqual(extractChain(result), expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testcase6(t *testing.T) {
	result := addTwoNumbers(buildChain([]int{0}), buildChain([]int{2, 7, 8}))
	expected := []int{2, 7, 8}

	if !reflect.DeepEqual(extractChain(result), expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}
