package main

import (
	"fmt"
	"errors"
)

type stack []byte

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s *stack) Push(element byte) {
	if s.Empty() {
		*s = []byte{element}
	} else {
		*s = append(*s, element)
	}
}

func (s *stack) Pop() (byte, error) {
	if s.Empty() {
		return '.', errors.New("empty stack")
	}
	top := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return top, nil
}

func (s *stack) Peek() (byte, error) {
	if s.Empty() {
		return '.', errors.New("empty stack")
	}
	top := (*s)[len(*s) - 1]
	return top, nil
}

func minAddToMakeValid(s string) int {
	var st stack
	inserts := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			st.Push(c)
		} else { // c = ')'
			if st.Empty() {
				inserts++
			} else {
				peek, err := st.Peek()
				if err == nil {
					if peek == '(' {
						_,_ = st.Pop()
					} else { // peel = ')'
						st.Push(c)
					}
				}
			}
		}
	}
	if st.Empty() {
		return inserts
	}
	return inserts + len(st)
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("()"))
	fmt.Println(minAddToMakeValid("()))(("))
}
