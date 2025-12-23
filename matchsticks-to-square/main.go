package main

import (
	"fmt"
	"sort"
)

// 1 stick
// 2 sick

// 3
// 1 1 2 2 3 4

// 1 1
// 1 2

func haveSumEqualTarget(target int, matchsticks []int, usedSticks map[int]bool) bool {
	fast := 0
	slow := 0
	temp := 0
	for fast >= slow && fast <= len(matchsticks) {

		if temp == target {
			return true
		} else if temp < target {
			if usedSticks[fast] {
				fast++
				continue
			}
			temp += matchsticks[fast]
			usedSticks[fast] = true
			fast++
		} else if temp > target {
			temp -= matchsticks[slow]
			usedSticks[slow] = false
			slow++
		}
	}
	return false
}

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	targetA := sum / 4
	if sum != targetA*4 {
		return false
	}
	square := make([]int, 4)

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	var dfs func(matchsticks []int, target, pos int, square []int) bool
	dfs = func(matchsticks []int, target, pos int, square []int) bool {
		if pos >= len(matchsticks) {
			return square[0] == target && square[1] == target && square[2] == target
		}

		for i := 0; i < 4; i++ {
			if square[i]+matchsticks[pos] > target {
				continue
			}
			square[i] += matchsticks[pos]
			if dfs(matchsticks, target, pos+1, square) {
				return true
			}
			square[i] -= matchsticks[pos]
		}
		return false
	}

	return dfs(matchsticks, targetA, 0, square)

}

func main() {
	fmt.Println(makesquare([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
	// haveSumEqualTarget(4, []int{3, 3, 3, 3}, map[int]bool{})
}
