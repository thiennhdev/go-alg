package main

import "fmt"

func moveZeros(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[slow] = nums[slow], nums[i]
			slow++
		}
	}
}

func main() {
	arr := []int{3, 0, 1, 0, 1, 3, 8, 9}
	moveZeros(arr)
	fmt.Println("arr: ", arr)
}
