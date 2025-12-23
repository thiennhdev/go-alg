package main

import "fmt"

func subarraySumLongest(nums []int, target int) int {
	maxLength := 0
	total := 0
	ptr := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		for total > target {
			total -= nums[ptr]
			ptr++
		}
		if i-ptr > maxLength {
			maxLength = i - ptr
		}
	}

	return maxLength + 1
}

func main() {
	fmt.Println(subarraySumLongest([]int{1, 6, 3, 1, 2, 4, 5}, 10))
}
