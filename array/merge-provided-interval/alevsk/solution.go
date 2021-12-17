package main

import (
	"fmt"
)

func mergeIntervals(intervals [][]int, newInterval []int) [][]int {
	// [[1,3], [7,10], [13,16]]
	// [6,11]
	//
	// ***...****..****
	//      ******
	//
	// ***...****..****
	//        **
	//
	// ***...****..****
	//      ****
	//
	// ***...****..****
	//      ********
	l := len(intervals)
	solution := [][]int{}
	nStart := newInterval[0]
	nEnd := newInterval[1]
	var inserted bool
	for i := 0; i < l; i++ {
		currentInterval := intervals[i]
		if currentInterval[1] < nStart {
			solution = append(solution, currentInterval)
		} else if currentInterval[0] > nEnd {
			if !inserted {
				solution = append(solution, []int{nStart, nEnd})
				inserted = true
			}
			solution = append(solution, currentInterval)
		} else {
			if currentInterval[0] < nStart {
				nStart = currentInterval[0]
			}
			if currentInterval[1] > nEnd {
				nEnd = currentInterval[1]
			}
		}
	}
	if !inserted {
		solution = append(solution, []int{nStart, nEnd})
		inserted = true
	}
	return solution

}

func main() {
	fmt.Println(mergeIntervals([][]int{{11,33}, {77,100}}, []int{5,200})) // [[5 200]]
	fmt.Println(mergeIntervals([][]int{{11,33}, {77,100}}, []int{5,9})) // [[5 9] [11 33] [77 100]]
	fmt.Println(mergeIntervals([][]int{{1,3}, {7,10}}, []int{11,13})) // [[1 3] [7 10] [11 13]]
	fmt.Println(mergeIntervals([][]int{{1,3}, {7,10}}, []int{2,7})) // [[1 10]]
	fmt.Println(mergeIntervals([][]int{{1,3}, {7,10}, {13,16}, {20,21}}, []int{8,14})) // [[1 3] [7 16] [20 21]]
}

