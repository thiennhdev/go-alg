package main

import (
	"fmt"
)

func findBoolean(arr []bool) int {

	left := 0
	right := len(arr) - 1
	boud := -1
	for right >= left {
		mid := (right + left) / 2
		if !arr[mid] {
			left = mid + 1
		} else {
			boud = mid
			right = mid - 1
		}
	}

	return boud
}

func main() {
	check := findBoolean([]bool{true, true, true, true, true})
	fmt.Println(check)
}
