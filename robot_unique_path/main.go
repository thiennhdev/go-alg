package main

import "fmt"

func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dp func(row, col int) int
	dp = func(row, col int) int {
		if row == m-1 && col == n-1 {
			return 1 // có 1 cách
		}

		if row < 0 || row >= m || col < 0 || col >= n {
			return 0 // out
		}

		if memo[row][col] != -1 {
			return memo[row][col]
		}

		// right
		right := dp(row, col+1)

		//down
		down := dp(row+1, col)

		res := down + right
		memo[row][col] = res
		return res

	}

	return dp(0, 0)
}

func uniquePathsIi(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(row, col int) int
	dp = func(row, col int) int {

		if row < 0 || row >= m || col < 0 || col >= n {
			return 0 // out
		}

		if obstacleGrid[row][col] == 1 {
			memo[row][col] = 0
			return 0
		}

		if row == m-1 && col == n-1 {
			return 1 // có 1 cách
		}

		if memo[row][col] != -1 {
			return memo[row][col]
		}

		// right
		right := dp(row, col+1)

		//down
		down := dp(row+1, col)

		res := down + right
		memo[row][col] = res
		return res

	}

	return dp(0, 0)

}

func main() {
	// fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePathsIi([][]int{
		{0, 0},
		{0, 1},
	}))
}
