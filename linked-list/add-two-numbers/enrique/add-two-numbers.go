package enrique

// ListNode is Definition for a singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	solution := &ListNode{}
	carry := false

	sum := solution
	currentL1 := l1
	currentL2 := l2
	for {
		if currentL1 == nil && currentL2 == nil && !carry{
			break
		}

		a, b, c := 0, 0, 0
		if currentL1 != nil {
			a = currentL1.Val
			currentL1 = currentL1.Next
		}

		if currentL2 != nil {
			b = currentL2.Val
			currentL2 = currentL2.Next
		}

		if carry {
			c = 1
		}

		total := a + b + c
		carry = total > 9
		sum.Val = total % 10

		if currentL2 != nil || currentL1 != nil || carry {
			sum.Next = &ListNode{}
			sum = sum.Next
		}
	}

	return solution
}
