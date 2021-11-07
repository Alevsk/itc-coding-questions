package enrique

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if  l1 == nil && l2 == nil {
		return nil
	}

	solutionRoot := &ListNode{}

	iterL1 := l1
	iterL2 := l2
	mergedSolutionIter := solutionRoot
	for iterL1 != nil && iterL2 != nil {
		if iterL1.Val < iterL2.Val {
			mergedSolutionIter.Val = iterL1.Val
			iterL1 = iterL1.Next
		} else {
			mergedSolutionIter.Val = iterL2.Val
			iterL2 = iterL2.Next
		}

		mergedSolutionIter.Next = &ListNode{}
		mergedSolutionIter = mergedSolutionIter.Next
	}

	if iterL1 != nil {
		mergedSolutionIter.Val = iterL1.Val
		mergedSolutionIter.Next = iterL1.Next
	}

	if iterL2 != nil {
		mergedSolutionIter.Val = iterL2.Val
		mergedSolutionIter.Next = iterL2.Next
	}

	return solutionRoot
}
