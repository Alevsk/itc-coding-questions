package main

import (
	"errors"
	"fmt"
)

type stack []rune

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s *stack) Push(r rune) {
	if s.Empty() {
		*s = []rune{r}
	} else {
		*s = append(*s, r)
	}
}

func (s *stack) Pop() (rune, error) {
	if s.Empty() {
		return '.', errors.New("empty stack")
	}
	topValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return topValue, nil
}

func isValid(s string) bool {
	var st stack
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			st.Push(r)
			break
		case ')':
			val, err := st.Pop()
			if err != nil {
				return false
			}
			if val != '(' {
				return false
			}
			break
		case '}':
			val, err := st.Pop()
			if err != nil {
				return false
			}
			if val != '{' {
				return false
			}
			break
		case ']':
			val, err := st.Pop()
			if err != nil {
				return false
			}
			if val != '[' {
				return false
			}
			break
		default:
			return false
		}
	}
	return st.Empty()
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
