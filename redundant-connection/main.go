package main

import "fmt"

// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := 0; i <= len(edges); i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			parent[rootY] = rootX
			return true
		}
		return false
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return []int{}

}

// use dsu
func main() {
	// fmt.Println(findRedundantConnection([][]int{
	// 	{1, 2},
	// 	{1, 3},
	// 	{2, 3},
	// }))

	// fmt.Println("2: ", findRedundantConnection([][]int{
	// 	{1, 2},
	// 	{2, 3},
	// 	{3, 4},
	// 	{1, 4},
	// 	{1, 5},
	// }))

	fmt.Println("3: ", findRedundantConnection([][]int{
		{1, 4},
		{3, 4},
		{1, 3},
		{1, 2},
		{4, 5},
	}))
}

// [[1,4],[3,4],[1,3],[1,2],[4,5]]
