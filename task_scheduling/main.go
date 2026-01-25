package main

import "fmt"

func taskScheduling(tasks []string, requirements [][]string) []string {
	graph := make(map[string][]string)
	indegree := make(map[string]int)

	for _, t := range tasks {
		graph[t] = []string{}
		indegree[t] = 0
	}

	for _, r := range requirements {
		from, to := r[0], r[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	queue := []string{}
	for _, t := range tasks { // 👈 dùng slice để giữ order ổn định
		if indegree[t] == 0 {
			queue = append(queue, t)
		}
	}

	result := []string{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)

		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(taskScheduling([]string{"a", "b", "c", "d"},
		[][]string{
			{"a", "b"},
			{"c", "b"},
			{"b", "d"},
		}))
}
