package main

import (
	"fmt"
)

func isPalindrome(s string, left, right int) bool {
	for {
		if left >= right {
			return true
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
}

func longestPalindrome(s string) string {
	max := 0
	sol := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			isPal := isPalindrome(s, i, j)

			if isPal {
				palLength := j - i + 1
				if palLength > max {
					max = palLength
					sol = s[i : j+1]
				}
			}
		}
	}
	return sol
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}
