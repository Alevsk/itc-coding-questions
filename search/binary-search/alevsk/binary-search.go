package main

import (
	"fmt"
)

func binarySearch(array []int, n int) bool {
	left := 0
	right := len(array) - 1
	for {
		if left > right {
			return false
		}
		mid := (left + right) / 2
		if n == array[mid] {
			return true
		} else if n > array[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func binarySearchRecursive(array []int, n, left, right int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	if n == array[mid] {
		return true
	} else if n > array[mid] {
		return binarySearchRecursive(array, n, mid+1, right)
	} else {
		return binarySearchRecursive(array, n, left, mid-1)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("recursive:\n")
	fmt.Println(binarySearchRecursive(numbers, 8, 0, len(numbers)-1))
	fmt.Println(binarySearchRecursive(numbers, 1, 0, len(numbers)-1))
	fmt.Println(binarySearchRecursive(numbers, 20, 0, len(numbers)-1))
	fmt.Println(binarySearchRecursive(numbers, 22, 0, len(numbers)-1))
	fmt.Println("")
	fmt.Println("iterative:\n")
	fmt.Println(binarySearch(numbers, 8))
	fmt.Println(binarySearch(numbers, 1))
	fmt.Println(binarySearch(numbers, 20))
	fmt.Println(binarySearch(numbers, 22))
}
