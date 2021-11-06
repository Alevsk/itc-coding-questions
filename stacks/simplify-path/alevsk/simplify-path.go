package main

import (
	"errors"
	"fmt"
	"strings"
)

type stack []string

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s *stack) Push(path string) {
	if s.Empty() {
		*s = []string{path}
	} else {
		*s = append(*s, path)
	}
}

func (s *stack) Pop() (string, error) {
	if s.Empty() {
		return "", errors.New("empty list")
	}
	topValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return topValue, nil
}

func (s *stack) First() (string, error) {
	if s.Empty() {
		return "", errors.New("empty list")
	}
	firstValue := (*s)[0]
	*s = (*s)[1:]
	return firstValue, nil
}

func simplifyPath(path string) string {
	aPath := strings.Split(path, "/")
	var s stack
	for _, subPath := range aPath {
		switch subPath {
		case "", ".":
			break
		case "..":
			_, _ = s.Pop()
			break
		default:
			s.Push(subPath)
		}
	}
	var solution string
	for {
		subPath, err := s.First()
		if err != nil {
			if len(solution) == 0 {
				solution = "/"
			}
			break
		}
		solution = solution + "/" + subPath
	}
	return solution
}

func main() {
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
