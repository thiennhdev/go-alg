package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func dfs(root *Node, depth int) int {
	if root == nil {
		return depth
	}

	left := dfs(root.left, depth+1)
	right := dfs(root.right, depth+1)
	max := left
	if right > max {
		max = right
	}
	return max
}

func treeMaxDepth(root *Node) int {
	return dfs(root, -1)
}

func main() {
	fmt.Println(treeMaxDepth(nil))
}
