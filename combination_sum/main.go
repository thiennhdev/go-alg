package main

import (
	"fmt"
	"sort"
)

// func exist(result [][]int, path []int) bool {
// 	sort.Ints(path)
// 	for _, r := range result {
// 		if len(r) != len(path) {
// 			continue
// 		}
// 		sort.Ints(r)
// 		length := len(r)
// 		check := 0
// 		for i := 0; i < length; i++ {
// 			if r[i] != path[i] {
// 				break
// 			}
// 			check++
// 		}
// 		if check == length {
// 			return true
// 		}
// 	}

// 	return false
// }

// func combinationSum(candidates []int, target int) [][]int {
// 	result := make([][]int, 0)
// 	var dfs func(curr int, path []int)
// 	dfs = func(curr int, path []int) {

// 		if curr == target {
// 			if !exist(result, path) {
// 				result = append(result, path)
// 			}
// 			return
// 		}

// 		for _, candidate := range candidates {
// 			if curr+candidate > target {
// 				continue
// 			}
// 			dfs(curr+candidate, append(path, candidate))
// 		}
// 	}

// 	dfs(0, []int{})
// 	return result
// }

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(nums []int, startIndex int, remaining int, path []int)
	dfs = func(nums []int, startIndex int, remaining int, path []int) {
		if remaining == 0 {
			// Append a copy of the path to the result
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i < len(nums); i++ {
			num := nums[i]
			if remaining-num < 0 {
				continue
			}
			dfs(nums, i, remaining-num, append(path, num))
		}
	}

	sort.Ints(candidates)
	dfs(candidates, 0, target, []int{})
	return res
}

func main() {
	fmt.Println("combinationSum: ", combinationSum([]int{2, 3, 6, 7}, 7))
}
