package main

import "fmt"

func getIsland(grid [][]byte, visited [][]bool, col, row int) {

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}

	for _, direction := range directions {
		newCol := col + direction[0]
		newRow := row + direction[1]

		if newCol >= 0 && newCol < len(grid) && newRow >= 0 && newRow < len(grid[0]) {
			if grid[newCol][newRow] == '1' && !visited[newCol][newRow] {
				visited[newCol][newRow] = true
				getIsland(grid, visited, newCol, newRow)
			}
		}
	}
}

func numIslands(grid [][]byte) int {

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	result := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' && !visited[row][col] {
				visited[row][col] = true
				getIsland(grid, visited, row, col)
				result++
			}
		}
	}

	return result
}

// func countNumberOfIslands(grid [][]int) int {
//     // WRITE YOUR BRILLIANT CODE HERE
//     return 0
// }

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))

}
