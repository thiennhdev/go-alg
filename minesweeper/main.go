package main

import "fmt"

// Input: board = [["E","E","E","E","E"],
// ["E","E","M","E","E"],
// ["E","E","E","E","E"],
// ["E","E","E","E","E"]],
// click = [3,0]

// Output: [["B","1","E","1","B"],["B","1","M","1","B"],["B","1","1","1","B"],["B","B","B","B","B"]]

func getNeighbor(maxRow, maxCol, row, col int) [][]int {
	positions := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	res := make([][]int, 0)
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

func updateBoard(board [][]byte, click []int) [][]byte {
	maxRow := len(board)
	maxCol := len(board[0])

	queue := make([][]int, 0)
	queue = append(queue, []int{
		click[0], click[1],
	})

	clickValue := board[click[0]][click[1]]
	switch clickValue {
	case 'M':
		board[click[0]][click[1]] = 'X'
		return board
	case 'E':
		break
	default:
		return board
	}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[i]

			if board[q[0]][q[1]] != 'E' {
				continue
			}
			board[q[0]][q[1]] = 'B'

			neighbors := getNeighbor(maxRow, maxCol, q[0], q[1])
			countBomb := 0
			tempQueue := make([][]int, 0)
			for _, neighbor := range neighbors {
				value := board[neighbor[0]][neighbor[1]]
				switch value {
				case 'M', 'X':
					countBomb++
				case 'E':
					tempQueue = append(tempQueue, []int{
						neighbor[0], neighbor[1],
					})
				}
			}

			if countBomb > 0 {
				board[q[0]][q[1]] = byte('0' + countBomb)
				continue
			}

			queue = append(queue, tempQueue...)
		}

		queue = queue[n:]
	}

	return board
}

func main() {
	// board := [][]byte{
	// 	{'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'M', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E'},
	// }

	board := [][]string{
		{"E", "E", "E", "E", "E"},
		{"E", "E", "M", "E", "E"},
		{"E", "E", "E", "E", "E"},
		{"E", "E", "E", "E", "E"},
	}
	click := []int{3, 0}

	fmt.Println(updateBoardString(board, click))
}
