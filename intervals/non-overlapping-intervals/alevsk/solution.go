package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var solution int
	l := len(intervals)
	prevInterval := intervals[0]
	for i := 1; i < l; i++ {
		currentInterval := intervals[i]
		if currentInterval[0] < prevInterval[1] {
			solution++
		} else {
			prevInterval = currentInterval
		}
	}
	return solution
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3},{3,4},{1,3}})) // 1
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{1,2},{1,2}})) // 2
	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3}})) // 0
	fmt.Println(eraseOverlapIntervals([][]int{{1,100},{11,22},{1,11},{2,12}}))// 2
}
