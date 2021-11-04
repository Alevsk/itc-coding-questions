package main

import "fmt"

func Solution(input string) string {
	var solution []string
	tmp := ""
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			if len(tmp) > 0 {
				solution = append(solution, tmp)
			}
			tmp = ""
		} else {
			tmp = fmt.Sprintf("%s%c", tmp, input[i])
		}
	}
	if len(tmp) > 0 {
		solution = append(solution, tmp)
	}
	var reversed string
	for i := len(solution) - 1; i >= 0; i-- {
		if i == 0 {
			reversed = fmt.Sprintf("%s%s", reversed, solution[i])
		} else {
			reversed = fmt.Sprintf("%s%s ", reversed, solution[i])
		}

	}
	return reversed
}

func main() {
	fmt.Println(Solution("the sky is blue"))
	fmt.Println(Solution("  hello world  "))
	fmt.Println(Solution(" good   example"))
	fmt.Println(Solution("  Bob    Loves  Alice   "))
	fmt.Println(Solution("Alice does not even like bob"))
}
