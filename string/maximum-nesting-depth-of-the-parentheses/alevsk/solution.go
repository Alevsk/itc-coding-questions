func maxDepth(s string) int {
    max := 0
    count := 0
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' {
            count++
            if count > max {
                max = count
            }
        } else if c == ')' {
            count--
        }
    }
    return max
}