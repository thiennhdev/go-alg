package main

import "fmt"

func getNeighbors(maxRow, maxCol, row, col int) [][]int {
	res := make([][]int, 0, 4)

	positions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, p := range positions {
		newRow, newCol := row+p[0], col+p[1]
		if newRow < 0 || newRow >= maxRow || newCol < 0 || newCol >= maxCol {
			continue
		}
		res = append(res, []int{newRow, newCol})
	}
	return res
}

func bfsFromEdges(heights [][]int, queue [][]int, visited [][]bool) {
	maxRow := len(heights)
	maxCol := len(heights[0])

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			r, c := node[0], node[1]
			curH := heights[r][c]

			neighbors := getNeighbors(maxRow, maxCol, r, c)
			for _, nb := range neighbors {
				nr, nc := nb[0], nb[1]
				if visited[nr][nc] {
					continue
				}

				if heights[nr][nc] >= curH {
					visited[nr][nc] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}
		queue = queue[n:]
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)
	if len(heights) == 0 || len(heights[0]) == 0 {
		return res
	}

	maxRow := len(heights)
	maxCol := len(heights[0])

	pVisited := make([][]bool, maxRow)
	aVisited := make([][]bool, maxRow)
	for i := 0; i < maxRow; i++ {
		pVisited[i] = make([]bool, maxCol)
		aVisited[i] = make([]bool, maxCol)
	}

	pacificQueue := make([][]int, 0, maxRow+maxCol)
	atlantisQueue := make([][]int, 0, maxRow+maxCol)

	for c := 0; c < maxCol; c++ {
		pacificQueue = append(pacificQueue, []int{0, c})
		pVisited[0][c] = true
	}
	for r := 0; r < maxRow; r++ {
		pacificQueue = append(pacificQueue, []int{r, 0})
		pVisited[r][0] = true
	}

	for c := 0; c < maxCol; c++ {
		atlantisQueue = append(atlantisQueue, []int{maxRow - 1, c})
		aVisited[maxRow-1][c] = true
	}
	for r := 0; r < maxRow; r++ {
		atlantisQueue = append(atlantisQueue, []int{r, maxCol - 1})
		aVisited[r][maxCol-1] = true
	}

	bfsFromEdges(heights, pacificQueue, pVisited)
	bfsFromEdges(heights, atlantisQueue, aVisited)

	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if pVisited[r][c] && aVisited[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

func main() {
	fmt.Println("pacificAtlantic:",
		pacificAtlantic([][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		}),
	)
}
