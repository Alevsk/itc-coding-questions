package main

import (
	"fmt"
)

func checkIfIsValidSubstr(s string, i, j int) bool {
	var observed [300]int
	for l := i; l <= j; l++ {
		index := s[l] - 'a'
		observed[index]++
		if observed[index] > 1 {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			isValidSubstr := checkIfIsValidSubstr(s, i, j)
			if isValidSubstr {
				substrLength := j - i + 1
				if substrLength > maxLength {
					maxLength = substrLength
				}
			} else {
				break
			}
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}
