package main

import (
	"fmt"
)

func min(a, b int) int {

	if b < a {
		return b
	}

	return a
}

func max(a, b int) int {

	if b > a {
		return b
	}

	return a
}

func calculateMinimumHP(dungeon [][]int) int {
	rowLen := len(dungeon)
	colLen := len(dungeon[0])

	memo := make([][]int, rowLen)
	vis := make([][]bool, rowLen)
	for i := 0; i < rowLen; i++ {
		memo[i] = make([]int, colLen)
		vis[i] = make([]bool, colLen)
	}

	var dp func(row, col int) int
	const INF = int(1e9)

	dp = func(row, col int) int {

		if row < 0 || row >= rowLen || col < 0 || col >= colLen {
			return INF
		}

		if vis[row][col] {
			return memo[row][col]
		}

		vis[row][col] = true

		if row == rowLen-1 && col == colLen-1 {
			memo[row][col] = max(1, 1-dungeon[row][col])
			return memo[row][col]
		}

		// right
		next := min(dp(row, col+1), dp(row+1, col))
		memo[row][col] = max(1, next-dungeon[row][col])
		return memo[row][col]

	}
	return dp(0, 0)
}

func main() {
	fmt.Println(calculateMinimumHP(
		[][]int{
			{-3},
			{-7},
		},
	))
}
