package main

import "math"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func dfs(tree *Node) (int, bool) {
	if tree == nil {
		return 0, true
	}

	heightL, checkL := dfs(tree.left)
	heightR, checkR := dfs(tree.right)

	if !checkL || !checkR {
		return 0, false
	}

	if math.Abs(float64(heightL-heightR)) > 1 {
		return 0, false
	}

	return int(math.Max(float64(heightL), float64(heightR))) + 1, true
}

func isBalanced(tree *Node) bool {
	_, balance := dfs(tree)
	return balance
}

func newNode(v int) *Node { return &Node{val: v} }

func buildExampleTree() *Node {
	n1 := newNode(1)
	n2 := newNode(2)
	n3 := newNode(3)
	n4 := newNode(4)
	n5 := newNode(5)
	n6 := newNode(6)
	n7 := newNode(7)
	n8 := newNode(8)

	n1.left, n1.right = n2, n3
	n2.left, n2.right = n4, n5
	n3.right = n6
	n4.left = n7
	n6.right = n8

	return n1
}

func main() {
	root := buildExampleTree()
	isBalanced(root)
}
