package main

import "fmt"

func buildGraph(vertices int, edges [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func dfsRecursive(graph map[int][]int, start int, visited map[int]bool) {
	visited[start] = true
	if _, ok := graph[start]; !ok {
		fmt.Println("Node not found")
		return
	}
	fmt.Println(start)

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited)
		}
	}
}

func dfs(graph map[int][]int, start int, visited map[int]bool) {
	stack := make([]int, 0)
	stack = append(stack, start)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[current] {
			continue
		}
		fmt.Println(current)
		visited[current] = true
		stack = append(stack, graph[current]...)
	}
}

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}}
	vertices := 3
	visited := make(map[int]bool)

	fmt.Println("--------------------------------")
	dfs(buildGraph(vertices, edges), 0, visited)
	fmt.Println("--------------Recursive------------------")
	visited = make(map[int]bool)
	dfsRecursive(buildGraph(vertices, edges), 0, visited)
	fmt.Println("--------------------------------")
}
