package main

import "fmt"

// Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
// Output: 2

// 2 -> [[1,1], [3,1]]
// 3 , 1(done) -> [[4,1]]

// 31 11
// 4 2

// 1 -1> 2
// 2 -3> 1
func networkDelayTime(times [][]int, n int, k int) int {
	max := 10000000
	graph := make(map[int][][]int)

	for _, time := range times {
		node := time[0]
		to := time[1]
		lenght := time[2]

		if _, ok := graph[node]; !ok {
			graph[node] = make([][]int, 0)
		}

		graph[node] = append(graph[node], []int{
			to, lenght,
		})
	}

	shortest := make(map[int]int)
	for i := 1; i <= n; i++ {
		shortest[i] = max // unknown
	}

	queue := make([]int, 0)
	queue = append(queue, k)
	shortest[k] = 0
	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			cur := queue[i]
			neighbors := graph[cur]
			for _, neighbor := range neighbors {
				nei := neighbor[0]
				value := neighbor[1]
				temp := value + shortest[cur]
				if temp < shortest[nei] {
					shortest[nei] = temp
				}
				queue = append(queue, nei)
			}
		}
		queue = queue[length:]
	}

	result := shortest[n]
	if result == max || result == 0 {
		return -1
	}
	return result
}

func main() {
	fmt.Println("networkDelayTime 1: ", networkDelayTime([][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}, 4, 2))

	fmt.Println("networkDelayTime 2: ", networkDelayTime([][]int{
		{1, 2, 1},
	}, 2, 2))
}
