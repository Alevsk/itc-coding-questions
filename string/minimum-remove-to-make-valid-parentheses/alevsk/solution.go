package main

import (
	"fmt"
)


func minRemoveToMakeValid(s string) string {
	open := 0
	sol := ""
	for _, c := range s {
		if c == '(' {
			open++
		} else if c == ')' {
			if open == 0 {
				continue
			}
			open--
		}
		sol += string(c)
	}
	if open == 0 {
		return sol
	} else {
		sol2 := ""
		index := len(sol) - 1
		for {
			if index < 0 {
				return sol2
			}
			if sol[index] == '(' && open > 0 {
				open--
			} else {
				sol2 = string(sol[index]) + sol2
			}
			index--
		}
	}
}

func main() {

	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)"))
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)))(("))
}
