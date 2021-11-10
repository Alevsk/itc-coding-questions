package main

import (
	"fmt"
	"unicode"
)

func IsAlphanumeric(r byte) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func isPalindrome(word string) bool {
	if len(word) == 0 {
		return true
	}
	left := 0
	right := len(word) - 1
	for {
		if left > right {
			break
		}
		if IsAlphanumeric(word[left]) && IsAlphanumeric(word[right]) {
			if unicode.ToLower(rune(word[left])) != unicode.ToLower(rune(word[right])) {
				return false
			}
			left++
			right--
		} else if !IsAlphanumeric(word[left]) {
			left++
		} else {
			right--
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("1a2"))
}
