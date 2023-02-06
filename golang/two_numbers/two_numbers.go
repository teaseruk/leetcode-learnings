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

	//	firstNode := &ListNode{Val: 0, Next: nil}
	result := l1
	lastL1 := result
	var tot int = 0
	l1Val := l1.Val
	l2Val := l2.Val

	for l1 != nil && l2 != nil {
		tot = tot + l1Val + l2Val

		l1.Val = tot % 10
		println(l1.Val)

		tot /= 10

		//prevNode.Next = &ListNode{Val: num, Next: nil}
		//prevNode = prevNode.Next

		lastL1 = l1
		if l1.Next != nil {
			l1 = l1.Next
			l1Val = l1.Val
		} else if l2.Next != nil {
				l1 = l2.Next
				l1Val = l1.Val
				l2 = nil
				l2Val = 0
			} else {
				l1Val = 0
				l2Val = 0
			}
	
		}

		if l2.Next != nil {
			l2 = l2.Next
			l2Val = l2.Val
		}
	}

	if tot > 0 {
		lastL1.Next = &ListNode{Val: tot, Next: nil}
	}

	return result
}
