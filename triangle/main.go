package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	memo := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, len(triangle[i]))
		vis[i] = make([]bool, len(triangle[i]))
	}

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r == n-1 { // last row
			return triangle[r][c]
		}
		if vis[r][c] {
			return memo[r][c]
		}
		down := dp(r+1, c)
		diag := dp(r+1, c+1)
		best := down
		if diag < best {
			best = diag
		}
		memo[r][c] = triangle[r][c] + best
		vis[r][c] = true
		return memo[r][c]
	}

	return dp(0, 0)
}
func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
