package main

import "fmt"

func removeDuplicates(arr []int) int {

	slow, fast := 0, 0
	for fast <= len(arr)-1 {
		for fast <= len(arr)-1 && arr[fast] == arr[slow] {
			fast++
		}
		if fast <= len(arr)-1 {
			slow++
			arr[slow] = arr[fast]
		}
	}
	arr = arr[:slow+1]
	return len(arr)
}

func main() {
	fmt.Println("removeDuplicates: ", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2}))
}
