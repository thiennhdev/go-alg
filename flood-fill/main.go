package main

import "fmt"

func getNeighbors(maxRow, maxCol, row, col int) [][]int {
	deltaRow := []int{-1, 0, 1, 0}
	deltaCol := []int{0, 1, 0, -1}

	res := make([][]int, 0, 4)
	for i := 0; i < 4; i++ {
		nr := row + deltaRow[i]
		nc := col + deltaCol[i]
		if nr >= 0 && nr < maxRow && nc >= 0 && nc < maxCol {
			res = append(res, []int{nr, nc})
		}
	}
	return res
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	numRow := len(image)
	numCol := len(image[0])
	startPixel := image[sr][sc]

	if startPixel == color {
		return image // Không cần làm gì nếu màu mới = màu cũ
	}

	visited := make([][]bool, numRow)
	for i := range visited {
		visited[i] = make([]bool, numCol)
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc})
	visited[sr][sc] = true

	head := 0
	for head < len(queue) {
		r, c := queue[head][0], queue[head][1]
		head++

		image[r][c] = color

		// BFS neighbors
		for _, nb := range getNeighbors(numRow, numCol, r, c) {
			nr, nc := nb[0], nb[1]

			if !visited[nr][nc] && image[nr][nc] == startPixel {
				visited[nr][nc] = true
				queue = append(queue, []int{nr, nc})
			}
		}
	}

	return image
}

func main() {
	arr := [][]int{
		{0, 1, 3, 4, 1},
		{3, 8, 8, 3, 3},
		{6, 7, 8, 8, 3},
		{12, 2, 8, 9, 1},
		{12, 3, 1, 3, 2},
	}

	fmt.Println(floodFill(arr, 2, 2, 9))
}
