package enrique

import (
	"strings"
)

const _whiteSpaceRune = 32

func reverseWords(s string) string {
	var wordsInString []string

	var currentWord strings.Builder
	for _, v := range s {
		if v == _whiteSpaceRune {
			if currentWord.Len() != 0 {
				wordsInString = append(wordsInString, currentWord.String())
				currentWord.Reset()
			}
			continue
		}

		currentWord.WriteRune(v)
	}

	if currentWord.Len() != 0 {
		wordsInString = append(wordsInString, currentWord.String())
	}

	var reverseWords strings.Builder
	for i := len(wordsInString) - 1; i >= 0; i-- {
		reverseWords.WriteString(wordsInString[i])
		if i != 0 {
			reverseWords.WriteString(" ")
		}
	}

	return reverseWords.String()
}

