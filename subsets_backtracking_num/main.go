package main

import "fmt"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(path []int, startIndex int)
	dfs = func(path []int, startIndex int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		for i := startIndex; i < len(nums); i++ {
			dfs(append(path, nums[i]), i+1)
		}
	}

	// sort.Ints(nums)
	dfs([]int{}, 0)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4, 5}))
}
