package main

import (
	"fmt"
)

func sell(input []int) int {
	if len(input) == 0 {
		return 0
	}
	lowest := input[0]
	maxP := 0
	for _, n := range input {
		if n < lowest {
			lowest = n
		}
		tmp := n - lowest
		if tmp > maxP {
			maxP = tmp
		}
	}
	return maxP
}

func main() {
	fmt.Println(sell([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(sell([]int{7, 6, 4, 3, 1}))    // 0
}
