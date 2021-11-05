package main

import "fmt"

func maxProduct(nums []int) int {
	leftP := make([]int, len(nums))
	rightP := make([]int, len(nums))
	tmpP := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			leftP[i] = 0
			tmpP = 1
		} else {
			tmpP = tmpP * nums[i]
			leftP[i] = tmpP
		}
	}
	tmpP = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			rightP[i] = 0
			tmpP = 1
		} else {
			tmpP = tmpP * nums[i]
			rightP[i] = tmpP
		}
	}

	maxProductLeft := leftP[0]
	maxProductRight := rightP[0]
	for i := 0; i < len(leftP); i++ {
		if leftP[i] > maxProductLeft {
			maxProductLeft = leftP[i]
		}
		if rightP[i] > maxProductRight {
			maxProductRight = rightP[i]
		}
	}

	if maxProductLeft > maxProductRight {
		return maxProductLeft
	}
	return maxProductRight
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
