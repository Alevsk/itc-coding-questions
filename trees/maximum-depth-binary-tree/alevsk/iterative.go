package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

type QueueNode struct {
	value *TreeNode
	next  *QueueNode
}

type Queue struct {
	head *QueueNode
}

func (q Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Put(node *TreeNode) {
	newQueueNode := &QueueNode{
		value: node,
	}
	if q.Empty() {
		q.head = newQueueNode
	} else {
		currentNode := q.head
		for {
			if currentNode.next == nil {
				currentNode.next = newQueueNode
				break
			}
			currentNode = currentNode.next
		}
	}
}

func (q *Queue) First() (*TreeNode, error) {
	if q.Empty() {
		return nil, errors.New("empty queue")
	}
	firstValue := q.head
	q.head = q.head.next
	return firstValue.value, nil
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q Queue
	q.Put(root)
	level := 1
	maxLevel := 1
	levels := map[string]int{}
	for {
		if q.Empty() {
			break
		}
		currentNode, _ := q.First()
		if val, ok := levels[fmt.Sprintf("%p", currentNode)]; ok {
			level = val
		} else {
			levels[fmt.Sprintf("%p", currentNode)] = level
		}
		if currentNode.left != nil {
			levels[fmt.Sprintf("%p", currentNode.left)] = level + 1
			q.Put(currentNode.left)
		}
		if currentNode.right != nil {
			levels[fmt.Sprintf("%p", currentNode.right)] = level + 1
			q.Put(currentNode.right)
		}
		if level > maxLevel {
			maxLevel = level
		}

	}
	return maxLevel
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
