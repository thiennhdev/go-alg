package main

import "fmt"

var Result = [][]int{
	{1, 2, 3},
	{4, 5, 0},
}

func flatten(pos [][]int) [6]int {
	var key [6]int
	k := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			key[k] = pos[i][j]
			k++
		}
	}
	return key
}

func clone2D(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int(nil), src[i]...)
	}
	return dst
}

func generateNewPos(cur [][]int, used map[[6]int]bool) [][][]int {
	res := make([][][]int, 0)

	positions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	zeroRow, zeroCol := -1, -1
	for i := 0; i < len(cur); i++ {
		for j := 0; j < len(cur[i]); j++ {
			if cur[i][j] == 0 {
				zeroRow, zeroCol = i, j
				break
			}
		}
	}

	for _, pos := range positions {
		newRow := zeroRow + pos[0]
		newCol := zeroCol + pos[1]

		if newRow < 0 || newRow > 1 || newCol < 0 || newCol > 2 {
			continue
		}

		next := clone2D(cur)

		next[zeroRow][zeroCol], next[newRow][newCol] =
			next[newRow][newCol], next[zeroRow][zeroCol]

		key := flatten(next)
		if used[key] {
			continue
		}
		used[key] = true

		res = append(res, next)
	}

	return res
}

func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func numSteps(initPos [][]int) int {
	used := map[[6]int]bool{}
	used[flatten(initPos)] = true
	queue := make([][][]int, 0)
	queue = append(queue, initPos)
	step := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[i]
			if equal2D(cur, Result) {
				return step
			}
			newCombo := generateNewPos(cur, used)
			queue = append(queue, newCombo...)
		}
		step++
		queue = queue[n:]
	}

	return -1
}

func main() {
	initPos := [][]int{
		{4, 1, 3},
		{2, 0, 5},
	}
	fmt.Println("numSteps: ", numSteps(initPos))
}
