package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

// Tìm các child có val lớn hơn tổ tiên của nó
func dfs(root *Node, max int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.val >= max {
		max = root.val
		count++
	}
	left := dfs(root.left, max)
	right := dfs(root.right, max)

	return left + right + count
}

func visibleTreeNode(root *Node) int {
	if root == nil {
		return 1
	}
	return dfs(root, root.val)
}

func main() {
	root := &Node{
		val: 5,
		left: &Node{
			val: 4,
			left: &Node{
				val: 3,
			},
			right: &Node{
				val: 8,
			},
		},
		right: &Node{
			val: 6,
		},
	}
	fmt.Println("visibleTreeNode(root):", visibleTreeNode(root))
}
