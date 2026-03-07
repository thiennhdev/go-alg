package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	var dp func(i int, memo []int) int

	dp = func(i int, memo []int) int {

		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}
		res := max(dp(i-1, memo), dp(i-2, memo)+nums[i])
		memo[i] = res
		return res
	}
	dp(len(nums)-1, memo)

	max := 0
	for _, d := range memo {
		if d > max {
			max = d
		}
	}
	return max
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
