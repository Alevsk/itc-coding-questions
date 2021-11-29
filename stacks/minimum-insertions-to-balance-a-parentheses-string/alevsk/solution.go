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

func minInsertions(input string) int {
	var s stack
	inserts := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == ')' {
			if s.Empty() {
				inserts++
				s.Push('(')
				s.Push(c)
			} else {
				peek, err := s.Peek()
				if err == nil {
					if peek == ')' {
						_, _ = s.Pop()
						_, _ = s.Pop()
					} else { // peek = '('
						s.Push(c)
					}
				}
			}
		} else { // c = '('
			if s.Empty() {
				s.Push(c)
			} else {
				peek, err := s.Peek()
				if err == nil {
					if peek == '(' {
						s.Push(c)
					} else { // peek = ')'
						inserts++
						_,_ = s.Pop()
						_,_ = s.Pop()
						s.Push(c)
					}
				}
			}
		}
	}
	if s.Empty() {
		return inserts
	} else {
		for {
			if s.Empty() {
				return inserts
			}
			c, _ := s.Pop()
			if c == '(' {
				inserts += 2
			} else { // c = ')'
				inserts++
				_,_ = s.Pop()
			}
		}
	}
}

func main() {
	fmt.Println(minInsertions("())"))
	fmt.Println(minInsertions("())(())))"))
	fmt.Println(minInsertions("))())("))
	fmt.Println(minInsertions("()))"))
}
