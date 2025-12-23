package main

import "fmt"

func findFirstOccurrence(arr []int, target int) int {
	startPtr, endPtr := 0, len(arr)-1
	bound := -1
	for startPtr <= endPtr {
		midPtr := (startPtr + endPtr) / 2
		if arr[midPtr] == target {
			bound = midPtr
			endPtr = midPtr - 1
		} else if arr[midPtr] > target {
			endPtr = midPtr - 1
		} else {
			startPtr = midPtr + 1
		}

	}
	return bound
}

func main() {
	fmt.Println(findFirstOccurrence([]int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}, 3))
}
