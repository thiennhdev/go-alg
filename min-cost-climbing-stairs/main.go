package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dp func(i int) int
	dp = func(i int) int {
		if i <= 1 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = min(dp(i-1)+cost[i-1], dp(i-2)+cost[i-2])
		return memo[i]
	}

	return dp(n)
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{
		10, 15, 20,
	}))
}
