package two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := l1
	var tot int = 0
	l1Val := l1.Val
	l2Val := l2.Val

	for l1 != nil {
		tot = tot + l1Val + l2Val

		l1.Val = tot % 10

		tot /= 10

		if l1.Next == nil {
			if l2 != nil && l2.Next != nil {
				l1.Next = l2.Next
				l1 = l2.Next
				l1Val = l1.Val
				l2 = nil
			} else {
				if tot > 0 {
					l1.Next = &ListNode{Val: 0, Next: nil}
					l1 = l1.Next
				} else {
					l1 = nil
				}
				l1Val = 0
			}
			l2Val = 0
		} else {
			l1 = l1.Next
			l1Val = l1.Val

			if l2 != nil && l2.Next != nil {
				l2 = l2.Next
				l2Val = l2.Val
			} else {
				l2Val = 0
			}
		}
	}

	return result
}
