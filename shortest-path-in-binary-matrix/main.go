package main

import "fmt"

// Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
// Output: 4

func getNeighbor(maxRow, maxCol, row, col int) [][]int {
	res := make([][]int, 0)
	positions := [][]int{
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, 0},
		{1, -1},
		{1, 1},
	}

	for _, position := range positions {
		newRow := row + position[0]
		newCol := col + position[1]
		if newRow < 0 || newRow >= maxRow || newCol < 0 || newCol >= maxCol {
			continue
		}

		res = append(res, []int{
			newRow, newCol,
		})
	}

	return res
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}

	maxRow, maxCol := len(grid), len(grid[0])
	queue := make([][]int, 0)
	queue = append(queue, []int{
		0, 0,
	})

	vistied := make([][]bool, maxRow)
	for i := 0; i < maxRow; i++ {
		vistied[i] = make([]bool, maxCol)
	}
	vistied[0][0] = true
	level := 0
	for len(queue) > 0 {
		n := len(queue)
		level++
		for i := 0; i < n; i++ {
			cur := queue[i]
			if cur[0] == maxRow-1 && cur[1] == maxCol-1 {
				return level
			}
			neighbors := getNeighbor(maxRow, maxCol, cur[0], cur[1])
			for _, neighbor := range neighbors {
				value := grid[neighbor[0]][neighbor[1]]
				if value != 0 {
					continue
				}

				if vistied[neighbor[0]][neighbor[1]] {
					continue
				}
				vistied[neighbor[0]][neighbor[1]] = true

				queue = append(queue, []int{
					neighbor[0], neighbor[1],
				})
			}
		}

		queue = queue[n:]
	}

	return -1
}

func main() {
	fmt.Println("shortestPathBinaryMatrix 4: ", shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))

	fmt.Println("shortestPathBinaryMatrix -1: ", shortestPathBinaryMatrix([][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))

	fmt.Println("shortestPathBinaryMatrix -1: ", shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}))

	fmt.Println("shortestPathBinaryMatrix 1: ", shortestPathBinaryMatrix([][]int{
		{0},
	}))
}
