package main

import (
	"fmt"
	"sort"
	"errors"
)

type stack [][]int

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s *stack) Push(interval []int) {
	if s.Empty() {
		*s = [][]int{interval}
	} else {
		*s = append(*s, interval)
	}
}

func (s *stack) Pop() ([]int, error) {
	if s.Empty() {
		return nil, errors.New("empty stack")
	}
	top := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return top, nil
}

func (s *stack) Peek() ([]int, error) {
	if s.Empty() {
		return nil, errors.New("empty stack")
	}
	top := (*s)[len(*s) - 1]
	return top, nil	
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
    	
	sort.Slice(intervals, func(i, j int) bool {
  		return intervals[i][0] < intervals[j][0]
	})
	var solution stack
	l := len(intervals)
	i := 0
	for {
		if i >= l {
			break
		}
		interval := intervals[i]
		if solution.Empty() {
			solution.Push(interval)
		} else {
			peekInterval, _ := solution.Peek()
			if interval[0] > peekInterval[1] {
				solution.Push(interval)
			} else {
				topInterval, _ := solution.Pop()
				newStart := interval[0]
				newEnd := interval[1]
				if topInterval[0] < newStart {
					newStart = topInterval[0]
				}
				if topInterval[1] > newEnd {
					newEnd = topInterval[1]
				}
				solution.Push([]int{newStart, newEnd})
			}
		}
		i++
	}

	return solution
}

func main() {
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}})) // [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1,4},{4,5}})) // [[1,5]]
	fmt.Println(merge([][]int{{1,4},{0,0}})) // [[0 0] [1 4]]
	fmt.Println(merge([][]int{{1,4},{0,2},{3,5}})) // [[0,5]]
	fmt.Println(merge([][]int{{2,3},{2,2},{3,3},{1,3},{5,7},{2,2},{4,6}})) // [[1,3],[4,7]]
}
