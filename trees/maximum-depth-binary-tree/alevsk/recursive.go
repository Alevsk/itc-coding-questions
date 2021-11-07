package main

import (
	"fmt"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + maxDepth(root.left)
	right := 1 + maxDepth(root.right)
	if left > right {
		return left
	}
	return right
}

func main() {

	test1 := &TreeNode{
		value: 3,
		left: &TreeNode{
			value: 9,
		},
		right: &TreeNode{
			value: 20,
			left: &TreeNode{
				value: 15,
			},
			right: &TreeNode{
				value: 7,
			},
		},
	}

	test2 := &TreeNode{
		value: 1,
		left:  nil,
		right: &TreeNode{
			value: 2,
		},
	}

	var test3 *TreeNode

	test4 := &TreeNode{
		value: 0,
	}

	test5 := &TreeNode{
		value: 1,
		left: &TreeNode{
			value: 2,
			left: &TreeNode{
				value: 4,
			},
		},
		right: &TreeNode{
			value: 3,
			right: &TreeNode{
				value: 5,
			},
		},
	}

	fmt.Println(maxDepth(test1)) // 3
	fmt.Println(maxDepth(test2)) // 2
	fmt.Println(maxDepth(test3)) // 0
	fmt.Println(maxDepth(test4)) // 1
	fmt.Println(maxDepth(test5)) // 3
}
