package main

import "fmt"

func tribonacci(n int) int {
	memo := make([]int, n+3)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	memo[0] = 0
	memo[1] = 1
	memo[2] = 1

	var dp func(n int, memo []int) int
	dp = func(n int, memo []int) int {
		if n < 0 {
			return 0
		}
		if memo[n] != -1 {
			return memo[n]
		}

		res := dp(n-1, memo) + dp(n-2, memo) + dp(n-3, memo)
		memo[n] = res
		return res
	}

	return dp(n, memo)
}

func main() {
	fmt.Println(tribonacci(3))
}
