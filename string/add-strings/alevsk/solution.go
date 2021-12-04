func addStrings(num1 string, num2 string) string {
    i := len(num1) - 1
    j := len(num2) - 1
    solution := ""
    overflow := 0
    for {
        if i < 0 && j < 0 {
            break
        }
        digit1 := 0
        digit2 := 0
        if i >= 0 {
            digit1, _ = strconv.Atoi(string(num1[i]))
            i--
        }
        if j >= 0 {
            digit2, _ = strconv.Atoi(string(num2[j]))
            j--
        }
        sum := digit1 + digit2 + overflow
        if sum >= 10 {
            solution = fmt.Sprintf("%d", (sum - 10)) + solution
            overflow = 1
        } else {
            solution = fmt.Sprintf("%d", sum) + solution
            overflow = 0
        }
    }
    if overflow > 0 {
        solution = "1" + solution
    }
    return solution
}