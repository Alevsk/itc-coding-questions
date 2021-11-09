package main

import (
	"fmt"
	"strings"
	"sort"
)

func isAnagram(s string, t string) bool {
	word1 := strings.Split(s, "")
	word2 := strings.Split(t, "")
	sort.Strings(word1)
	sort.Strings(word2)
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
