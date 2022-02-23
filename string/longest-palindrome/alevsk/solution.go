func longestPalindrome(s string) int {
    lookup := map[rune]int{}
    for _, c := range s {
        lookup[c]++
    }
    maxLength := 0
    singleChar := 0
    for _, sum := range lookup {
        if sum % 2 == 0 {
            maxLength += sum
        } else {
            maxLength += sum - 1
            singleChar = 1
        }
    }
    
    return maxLength + singleChar
}