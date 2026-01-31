package main

import "fmt"

// Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
// Output: false
func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph)) // 0: uncolored, 1 / -1: two colors

	for start := 0; start < len(graph); start++ {
		if color[start] != 0 {
			continue
		}

		queue := make([]int, 0)
		queue = append(queue, start)
		color[start] = 1 // init color for this component

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if color[neighbor] == 0 {
					color[neighbor] = -color[node] // opposite color
					queue = append(queue, neighbor)
				} else if color[neighbor] == color[node] {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	fmt.Println("isBipartite: ", isBipartite([][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	}))
}
