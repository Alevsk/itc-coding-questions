package main

import "fmt"

func Reverse(word *[]rune, start, end int) {
	for {
		if start >= end {
			break
		}
		tmp := (*word)[start]
		(*word)[start] = (*word)[end]
		(*word)[end] = tmp
		start++
		end--
	}
}

func Solution(word []rune) []rune {
	Reverse(&word, 0, len(word)-1)
	index := 0
	wordStart := 0
	inWord := false
	for {
		if index >= len(word) {
			if inWord {
				Reverse(&word, wordStart, index-1)
			}
			break
		}
		if word[index] != ' ' && !inWord {
			wordStart = index
			inWord = true
		} else if word[index] == ' ' && inWord {
			Reverse(&word, wordStart, index-1)
			inWord = false
		}
		index++
	}
	return word
}

func main() {
	fmt.Println(string(Solution([]rune("LENIN IS FUNNY"))))
	fmt.Println(string(Solution([]rune("LA PALABRA MATA LA COSA"))))
	fmt.Println(string(Solution([]rune("LATE STAGE CAPITALISM"))))
	fmt.Println(string(Solution([]rune(""))))
	fmt.Println(string(Solution([]rune("PODCAST"))))
	fmt.Println(string(Solution([]rune("I"))))
	fmt.Println(string(Solution([]rune("  Bob    Loves  Alice   "))))
}
