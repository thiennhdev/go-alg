package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func binaryTreeRightSideView(root *Node) []int {
	res := make([]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		lastElement := queue[len(queue)-1]
		res = append(res, lastElement.val)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.left != nil {
				queue = append(queue, cur.left)
			}

			if cur.right != nil {
				queue = append(queue, cur.right)
			}
		}
	}
	return res
}

func main() {
	n8 := &Node{val: 8}
	n7 := &Node{val: 7}

	// Level 2
	n4 := &Node{val: 4, left: n7}
	n5 := &Node{val: 5, left: n8}
	n6 := &Node{val: 6}

	// Level 1
	n2 := &Node{val: 2, left: n4, right: n5}
	n3 := &Node{val: 3, right: n6}

	// Level 0 (root)
	root := &Node{val: 1, left: n2, right: n3}

	fmt.Println(binaryTreeRightSideView(root))
}
