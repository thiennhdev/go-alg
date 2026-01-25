package main

import "fmt"

// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.

// Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
// Output: 4
// chưa tối ưu lắm, tìm tất cả quả hư, đếm quả tốt, cuối cùng nếu fresh > 0 thì -1
// add quả tốt bị lan chưa duyệt vào queue

func getNeighbor(grid [][]int, maxRow, maxCol, row, col int, visited [][]bool) [][]int {
	positions := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	res := make([][]int, 0)

	for _, position := range positions {
		newRow := row + position[0]
		newCol := col + position[1]

		if newRow < 0 || newRow >= maxRow || newCol < 0 || newCol >= maxCol {
			continue
		}

		if visited[newRow][newCol] {
			continue
		}

		res = append(res, []int{
			newRow, newCol,
		})
	}

	return res
}

func orangesRotting(grid [][]int) int {
	maxRow := len(grid)
	maxCol := len(grid[0])

	res := 0

	visited := make([][]bool, maxRow)
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, maxCol)
	}
	queue := make([][]int, 0)
	// [(row, col), (row,col)]

	// find first roten
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			if grid[row][col] == 2 {
				queue = append(queue, []int{
					row, col,
				})
				visited[row][col] = true
			}
		}
	}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[i]
			neighbors := getNeighbor(grid, maxRow, maxCol, q[0], q[1], visited)

			for _, neighbor := range neighbors {
				if grid[neighbor[0]][neighbor[1]] == 0 {
					continue
				}

				grid[neighbor[0]][neighbor[1]] = 2

				visited[neighbor[0]][neighbor[1]] = true
				queue = append(queue, []int{
					neighbor[0], neighbor[1],
				})
			}
		}

		queue = queue[n:]
		if len(queue) > 0 {
			res++
		}
	}

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			if grid[row][col] == 1 {
				return -1
			}
		}
	}

	return res
}

func main() {
	fmt.Println(orangesRotting(
		[][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		},
	))

	fmt.Println(orangesRotting(
		[][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		},
	))
}
