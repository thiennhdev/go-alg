package main

import "fmt"

func subarraySumFixed(nums []int, k int) int {
	total := 0
	if len(nums) < 3 {
		return total
	}
	for i := 0; i < k; i++ {
		total += nums[i]
	}
	ptr := 0
	for i := k; i < len(nums); i++ {
		tempTotal := total + nums[i] - nums[ptr]
		if tempTotal > total {
			total = tempTotal
		}
		ptr++
	}
	return total
}

func main() {
	fmt.Println("subarraySumFixed: ", subarraySumFixed([]int{1, 2, 3, 7, 4, 1}, 3))
}
