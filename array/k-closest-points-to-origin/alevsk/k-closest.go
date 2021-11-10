package main

import (
	"fmt"
	"math"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	distances := map[float64][][]int{}
	var dToSort []float64
	for _, point := range points {
		distance := math.Sqrt(float64(point[0]*point[0]) + float64(point[1]*point[1]))

		if val, ok := distances[distance]; ok {
			distances[distance] = append(val, point)
		} else {
			distances[distance] = [][]int{point}
		}
		dToSort = append(dToSort, distance)
	}
	sort.Float64s(dToSort)
	var solution [][]int
	for i := 0; i < len(dToSort); i++ {
		mappedPoints := distances[dToSort[i]] //[][]int
		solution = append(solution, mappedPoints...)
		if len(solution) == k {
			break
		}
	}
	return solution
}

func main() {
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
