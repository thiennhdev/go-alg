package main

import (
	"fmt"
	"sort"
)

// Input: candidates = [10,1,2,7,6,1,5], target = 8
// 1 1 2 5 6 7 10

// 8 -1
// 7 -1
// 6 -2 = 4
// 4 -5 = -1 return
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)

	var backtracking func(pos int, path []int, cur int)
	backtracking = func(pos int, path []int, cur int) {
		if cur == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := pos; i < len(candidates); i++ {
			if cur-candidates[i] < 0 {
				return
			}
			if i > pos && candidates[i] == candidates[i-1] {
				continue
			}
			backtracking(i+1, append(path, candidates[i]), cur-candidates[i])

		}
	}
	backtracking(0, []int{}, target)

	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// nice
