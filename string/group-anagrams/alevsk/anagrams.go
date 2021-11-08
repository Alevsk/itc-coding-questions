package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	observed := map[string][]string{}
	var groups [][]string
	for _, str := range strs {
		wordsArray := strings.Split(str, "")
		sort.Strings(wordsArray)
		key := strings.Join(wordsArray, "")
		if _, ok := observed[key]; ok {
			observed[key] = append(observed[key], str)
		} else {
			observed[key] = []string{str}
		}
	}
	for _, anagrams := range observed {
		groups = append(groups, anagrams)
	}
	return groups
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(words))
}
