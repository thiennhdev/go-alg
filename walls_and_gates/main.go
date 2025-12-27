package main

import (
	"fmt"
)

const INF = 2147483647

type direction struct {
	row int
	col int
}

func bfs(dungeonMap [][]int, row, col, step, lenRow, lenCol int) {
	if row >= lenRow || col >= lenCol || row < 0 || col < 0 {
		return
	}

	if dungeonMap[row][col] == -1 {
		return
	}

	if dungeonMap[row][col] < step {
		return
	}

	if dungeonMap[row][col] > step {
		dungeonMap[row][col] = step
	}

	directions := []direction{
		{row: -1, col: 0},
		{row: 0, col: -1},
		{row: 1, col: 0},
		{row: 0, col: 1},
	}

	for _, d := range directions {
		bfs(dungeonMap, row+d.row, col+d.col, step+1, lenRow, lenCol)
	}
}

func mapGateDistances(dungeonMap [][]int) [][]int {
	rowLenght := len(dungeonMap)
	colLenght := len(dungeonMap[0])

	for row := 0; row < rowLenght; row++ {
		for col := 0; col < colLenght; col++ {
			if dungeonMap[row][col] != 0 {
				continue
			}

			// BFS call recursion
			bfs(dungeonMap, row, col, 0, rowLenght, colLenght)
		}
	}
	return dungeonMap
}

func main() {
	dungeonMap := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}

	fmt.Println(mapGateDistances(dungeonMap))
}
