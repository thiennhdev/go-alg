package main

import (
	"fmt"
	"slices"
)

// Input: orginal: [1,2,3], seqs: [[1,2], [1,3], [2,3]]
// DONT suggestion code
func sequenceReconstruction(original []int, seqs [][]int) bool {
	graph := make(map[int][]int)
	for _, seq := range seqs {
		for i := 1; i < len(seq); i++ {
			graph[seq[i-1]] = append(graph[seq[i-1]], seq[i])
		}
	}

	indegree := make(map[int]int)
	for node, neighbors := range graph {
		if _, ok := indegree[node]; !ok {
			indegree[node] = 0
		}
		for _, neighbor := range neighbors {
			indegree[neighbor] = indegree[neighbor] + 1
		}
	}
	fmt.Println("indegree: ", indegree)
	fmt.Println("graph: ", graph)

	queue := make([]int, 0)
	for node, neighbor := range indegree {
		if neighbor == 0 {
			queue = append(queue, node)
		}
	}
	result := make([]int, 0)
	fmt.Println("queue: ", queue)
	for len(queue) > 0 {
		cur := queue[0]
		if len(queue) > 1 {
			return false
		}
		queue = queue[1:]
		result = append(result, cur)
		for _, node := range graph[cur] {
			indegree[node] = indegree[node] - 1
			if indegree[node] == 0 {
				queue = append(queue, node)
			}
		}
	}

	fmt.Println("result: ", result)
	return slices.Equal(result, original)
}

func main() {
	fmt.Println("sequenceReconstruction: ", sequenceReconstruction(
		[]int{1, 2, 3},
		[][]int{
			{1, 2},
			{1, 3},
			// {2, 3},
		},
	))
}
