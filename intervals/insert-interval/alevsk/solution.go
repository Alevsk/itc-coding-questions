package main

import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
    	newStart := newInterval[0]
    	newEnd := newInterval[1]
    	var solution [][]int
	var inserted bool
	for _, interval := range intervals { 
		if interval[1] < newStart {
			solution = append(solution, interval)
		}  else if interval[0] > newEnd {
			if !inserted {
				solution = append(solution, []int{newStart, newEnd})
				inserted = true
			}
			solution = append(solution, interval)
		} else {
			if newStart >= interval[0] {
				newStart = interval[0]
			}
			if interval[1] >= newEnd {
				newEnd = interval[1]
			}
		}
	}
	if !inserted {
		solution = append(solution, []int{newStart, newEnd})
		inserted = true
	}
	return solution
}

func main() {
	fmt.Println(insert([][]int{{1,3},{6,9}}, []int{2,5})) // [[1,5],[6,9]]
	fmt.Println(insert([][]int{{1,2},{3,5},{6,7},{8,10}, {12,16}}, []int{4,8})) // [[1,2],[3,10],[12,16]]
	fmt.Println(insert([][]int{{2,5},{6,7},{8,9}}, []int{0,1})) // [[0 1] [2 5] [6 7] [8 9]]
}
