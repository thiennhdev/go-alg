package main

import "fmt"

func firstNotSmaller(arr []int, target int) int {
	startPtr, endPtr := 0, len(arr)-1
	midPtr := 0
	bound := -1
	for startPtr <= endPtr {
		midPtr = (endPtr + startPtr) / 2
		// feasible
		if arr[midPtr] >= target {
			bound = midPtr
			endPtr = midPtr - 1
		} else {
			startPtr = midPtr + 1
		}
	}
	return bound
}

func main() {
	fmt.Println(firstNotSmaller([]int{1, 3, 3, 5, 8, 8, 10}, 2))
}
