package main

import "fmt"

func containerWithMostWater(arr []int) int {
	startPtr, endPtr := 0, len(arr)-1
	area := 0
	for endPtr > startPtr {
		minHeight := arr[endPtr]
		if minHeight > arr[startPtr] {
			minHeight = arr[startPtr]
		}
		newArea := (endPtr - startPtr) * minHeight
		if newArea > area {
			area = newArea
		}

		if arr[startPtr] > arr[endPtr] {
			endPtr--
		} else {
			startPtr++
		}

	}

	return area
}

func main() {
	fmt.Println("containerWithMostWater: ", containerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
