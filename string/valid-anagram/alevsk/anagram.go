package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    observed := map[rune]int{}
    for _, w := range s {
        if _, ok := observed[w]; ok {
            observed[w] = observed[w] + 1
        } else {
            observed[w] = 1
        }
    }
    for _, w := range t {
        if val, ok := observed[w]; ok {
            if val <= 0 {
                return false
            }
            observed[w] = observed[w] - 1
        } else {
            return false
        }
    }
    return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
