package main

import "fmt"

// đfs duyet chieu sau => di sau vao 1 con
// bfs duyet theo chieu rong => lay het cac con lien quan
// 0 - 0 duyet bfs

func getNeighbors(row, col int) [][]int {
	res := make([][]int, 0)
	detalRow := []int{-2, -2, -1, 1, 2, 2, 1, -1}
	deltaCol := []int{-1, 1, 2, 2, 1, -1, -2, -2}
	for i := range len(detalRow) {
		r := row + detalRow[i]
		c := col + deltaCol[i]

		res = append(res, []int{r, c})
	}
	return res
}

func getKnightShortestPath(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	}

	// 0 - 0 duyet bfs
	// queue luu tru [x, y, steps]
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0, 0})

	visited := make(map[string]bool)
	visited["0-0"] = true

	i := 0
	for i < len(queue) {
		cur := queue[i]
		i++

		currX, currY, steps := cur[0], cur[1], cur[2]

		for _, neighbor := range getNeighbors(currX, currY) {
			nx, ny := neighbor[0], neighbor[1]

			if nx == x && ny == y {
				return steps + 1
			}

			key := fmt.Sprintf("%d-%d", nx, ny)
			if !visited[key] {
				visited[key] = true
				queue = append(queue, []int{nx, ny, steps + 1})
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(getKnightShortestPath(1, 1))
}
