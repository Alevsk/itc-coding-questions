package main

import (
	"fmt"
)

func isBadRevisionApiMock(n int) bool {
	versions := []bool{false, false, false, true, true, true}
	return versions[n-1]
}

func badRevision(n int) int {
	bad := n
	left := 1
	right := n
	for {
		if left > right {
			return bad
		}
		mid := (left + right) / 2
		if isBadRevisionApiMock(mid) {
			if mid < bad {
				bad = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}


func main() {
	fmt.Println(badRevision(6)) // 4
}
