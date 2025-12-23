package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	mapTemp := make(map[int]int)
	for i, num := range nums {
		temp := target - num
		if _, ok := mapTemp[temp]; ok {
			return []int{temp, num}
		}
		mapTemp[num] = i
	}
	return []int{}
}

func deduplication(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		res := []int{num}
		res = append(res, twoSum(nums[i+1:], target-num)...)
		if len(res) == 3 {
			result = append(result, res)
		}
	}

	return result
}

func main() {
	fmt.Println(deduplication([]int{3, 1, 1, 2}, 6))
}
