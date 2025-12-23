package main

import "fmt"

func main() {
	arr := []int{0, 10, 0}
	fmt.Println(peakOfMountainArray(arr))
}

func peakOfMountainArray(arr []int) int {
	startPtr, endPtr := 0, len(arr)-1
	bound := -1

	for endPtr >= startPtr {
		midPtr := (startPtr + endPtr) / 2
		if midPtr+1 <= len(arr)-1 && arr[midPtr] > arr[midPtr+1] {
			if bound == -1 || arr[bound] < arr[midPtr] {
				bound = midPtr
			}
			endPtr = midPtr - 1
		} else {
			startPtr = midPtr + 1
		}
	}
	return bound
}
