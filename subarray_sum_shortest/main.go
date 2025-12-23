package main

import (
	"fmt"
	"math"
)

func subarraySumShortest(nums []int, target int) int {
	slow := 0
	total := 0
	count := math.MaxInt
	for i, fastValue := range nums {
		total += fastValue
		for total > target {
			if total-nums[slow] < target {
				break
			}
			total -= nums[slow]
			slow++
		}

		if total >= target && count > i-slow+1 {
			count = i - slow + 1
		}
	}
	return count
}

func main() {
	fmt.Println(subarraySumShortest([]int{1, 4, 1, 7, 3, 0, 2, 5}, 10))
}
