package main

import "fmt"

func climbStairs(n int) int64 {
	memo := make([]int64, n+2)

	// init
	memo[1] = 1
	memo[2] = 2

	var dp func(n int, memo []int64) int64
	dp = func(n int, memo []int64) int64 {
		if memo[n] != 0 {
			return memo[n]
		}

		res := dp(n-1, memo) + dp(n-2, memo)

		memo[n] = res
		return res
	}

	return dp(n, memo)
}

func main() {
	fmt.Println("climbStairs: ", climbStairs(3))
}
