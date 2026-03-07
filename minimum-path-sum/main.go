package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	memo := make([][]int, rowLen)
	for i := 0; i < rowLen; i++ {
		memo[i] = make([]int, colLen)
		for j := 0; j < colLen; j++ {
			memo[i][j] = -1
		}
	}
	const INF = math.MaxInt
	var dp func(row, col int) int
	dp = func(row, col int) int {
		if row < 0 || col < 0 || row >= rowLen || col >= colLen {
			return INF
		}

		// end
		if row == rowLen-1 && col == colLen-1 {
			return grid[row][col]
		}

		if memo[row][col] != -1 {
			return memo[row][col]
		}

		right := dp(row, col+1)

		down := dp(row+1, col)

		if right < down {
			memo[row][col] = grid[row][col] + right
		} else {
			memo[row][col] = grid[row][col] + down
		}

		return memo[row][col]

	}

	return dp(0, 0)
}

func main() {
	// fmt.Println(minPathSum([][]int{
	// 	{1, 3, 1},
	// 	{1, 5, 1},
	// 	{4, 2, 1},
	// }))

	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}
