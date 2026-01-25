package main

import "fmt"

// count_parents
func countParents(graph map[string][]string) map[string]int {
	counts := make(map[string]int)
	for node := range graph {
		counts[node] = 0
	}
	for parent := range graph {
		for _, node := range graph[parent] {
			counts[node]++
		}
	}
	return counts
}

// topo_sort
func topoSort(graph map[string][]string, taskTimes map[string]int) int {
	ans := 0
	queue := make([]string, 0)

	// distance (finish time)
	dis := make(map[string]int)

	counts := countParents(graph)

	// init
	for node := range counts {
		dis[node] = 0
		if counts[node] == 0 {
			queue = append(queue, node)
			dis[node] = taskTimes[node]
			if dis[node] > ans {
				ans = dis[node]
			}
		}
	}

	// topo + DP
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, child := range graph[node] {
			counts[child]--

			// relax longest path
			if dis[node]+taskTimes[child] > dis[child] {
				dis[child] = dis[node] + taskTimes[child]
			}

			if dis[child] > ans {
				ans = dis[child]
			}

			if counts[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	return ans
}

// task_scheduling_2
func taskScheduling2(tasks []string, times []int, requirements [][]string) int {
	graph := make(map[string][]string)
	taskTimes := make(map[string]int)

	for i, task := range tasks {
		graph[task] = []string{}
		taskTimes[task] = times[i]
	}

	for _, req := range requirements {
		from, to := req[0], req[1]
		graph[from] = append(graph[from], to)
	}

	return topoSort(graph, taskTimes)
}

func main() {
	tasks := []string{"abbreviate", "bricks", "cardinals", "dextrous", "fibre", "green", "height"}
	times := []int{1, 1, 1, 1, 1, 100, 1}
	req := [][]string{
		{"abbreviate", "bricks"},
		{"cardinals", "bricks"},
		{"dextrous", "bricks"},
		{"bricks", "fibre"},
		{"green", "fibre"},
	}

	fmt.Println(taskScheduling2(tasks, times, req)) // 101
}
