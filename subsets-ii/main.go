package main

import (
	"fmt"
	"sort"
)

// Example 1:

// Input: nums = [1,2,2]
// Output: [[]?? ,[1],[1,2],[1,2,2],[2],[2,2]]

// [] ?? check
// 1
// 1 2
// 1 2 2

// must no contain

func subsetsWithDup(nums []int) [][]int {

	result := make([][]int, 0)
	var dfs func(pos int, path []int)
	sort.Ints(nums)
	dfs = func(pos int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		// statement stop
		if len(path) == len(nums) {
			return
		}

		for i := pos; i < len(nums); i++ {

			if i > pos && nums[i] == nums[i-1] { // sort so can avoid duplicate
				continue
			}
			dfs(i+1, append(path, nums[i]))
		}
	}

	dfs(0, []int{})
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
