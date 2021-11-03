package main

import "fmt"

func twoSum(nums []int, target int) []int {
    observed := map[int]int{}
    for i, val := range nums {
        n := target - val
        if index, exist := observed[n]; exist {
            return []int{i, index} 
        } else {
            observed[val] = i
        }
    }
    return []int{}
}

func main() {
    fmt.Println(twoSum([]int{2,7,11,15}, 9))
}