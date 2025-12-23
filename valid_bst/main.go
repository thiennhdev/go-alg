package main

import "math"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func valid(root *Node, min, max int) bool {
	if root == nil {
		return true
	}

	if !(root.val > min && root.val < max) {
		return false
	}

	return valid(root.left, min, root.val) && valid(root.right, root.val, max)
}

func validBst(root *Node) bool {
	min, max := math.MinInt, math.MaxInt
	return valid(root, min, max)
}

func main() {

}
