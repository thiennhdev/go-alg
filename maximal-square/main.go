package main

import "fmt"

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])

	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	maxSide := 0

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r < 0 || c < 0 {
			return 0
		}
		if memo[r][c] != -1 {
			return memo[r][c]
		}

		if matrix[r][c] == '0' {
			memo[r][c] = 0
			return 0
		}

		val := 1 + min3(dp(r-1, c), dp(r, c-1), dp(r-1, c-1))
		memo[r][c] = val

		if val > maxSide {
			maxSide = val
		}
		return val
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dp(r, c)
		}
	}

	return maxSide * maxSide
}

func main() {
	matrix := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	fmt.Println(maximalSquare(matrix)) // 4 (square 2x2)
}
