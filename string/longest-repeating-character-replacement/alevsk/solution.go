package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	sol := 0
	for i := 0; i < len(s); i++ {
		var observed [30]int
		max := 0
		for j := i; j < len(s); j++ {
			index := s[j] - 'A'
			observed[index]++
			if observed[index] > max {
				max = observed[index]
			}
			length := j - i + 1
			//fmt.Println("length", length)
			//fmt.Println("max", max)
			tmpMax := max + k
			if tmpMax >= length {
				if length > sol {
					sol = length
				}
			}
		}
	}
	return sol
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1)) // 4
	fmt.Println(characterReplacement("AAAA", 2))    // 4
	fmt.Println(characterReplacement("ABBB", 2))    // 4
}
