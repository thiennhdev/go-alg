package main

import "fmt"

// Input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
// Output: 4

func maximalNetworkRank(n int, roads [][]int) int {
	result := 0
	indegree := make([]int, n)
	direct := make([][]bool, n)
	for i := 0; i < n; i++ {
		direct[i] = make([]bool, n)
	}

	for _, road := range roads {
		indegree[road[0]]++
		indegree[road[1]]++
		direct[road[0]][road[1]] = true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// pair (i j)
			isSame := 0
			if direct[i][j] || direct[j][i] {
				isSame = 1
			}
			temp := indegree[i] + indegree[j] - isSame
			if temp > result {
				result = temp
			}
		}
	}

	return result
}

func main() {
	fmt.Println(maximalNetworkRank(2, [][]int{
		{1, 0},
	}))
}
