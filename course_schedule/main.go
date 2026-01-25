package main

import "fmt"

// [0, 1] means you need to take course 1 before taking course 0.
// Input: n = 2, prerequisites = [[0, 1]]

// Output: true

// Explanation: we can take 1 first and then 0.

// Example 2:

// Input: n = 2, prerequisites = [[0, 1], [1, 0]]
func isValidCourseSchedule(n int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	indegree := make([]int, n)

	for _, p := range prerequisites {
		to, from := p[0], p[1] // from -> to
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		visited++

		for _, nei := range graph[cur] {
			indegree[nei]--
			if indegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	return visited == n
}

func main() {
	fmt.Println(isValidCourseSchedule(6, [][]int{
		{0, 1},
		{2, 3},
		{4, 5},
		{1, 2},
	}))
}
