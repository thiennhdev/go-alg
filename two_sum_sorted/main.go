package main

import "fmt"

func twoSumSorted(arr []int, target int) []int {
	startPtr := 0
	endPtr := len(arr) - 1
	for endPtr > startPtr {
		total := arr[endPtr] + arr[startPtr]
		if total > target {
			endPtr--
		} else if total < target {
			startPtr++
		} else {
			break
		}
	}
	return []int{startPtr, endPtr}
}

func main() {
	fmt.Println("twoSumSorted: ", twoSumSorted([]int{2, 3, 5, 8, 11, 15}, 5))
}
