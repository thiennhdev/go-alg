package main

// graph = [[1, 2], [0, 2, 3], [0, 1], [1]]
// A = 0
// B = 3

func shortestPath(graph [][]int, a int, b int) int {
	if a == b {
		return 0
	}

	n := len(graph)
	if a < 0 || a >= n || b < 0 || b >= n {
		return -1
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}

	queue := make([]int, 0)
	queue = append(queue, a) // start by a
	dist[a] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == b {
			return dist[cur]
		}
		for _, nxt := range graph[cur] {
			if dist[nxt] == -1 {
				dist[nxt] = dist[cur] + 1
				queue = append(queue, nxt)
			}
		}

	}

	return -1
}

func main() {

}
