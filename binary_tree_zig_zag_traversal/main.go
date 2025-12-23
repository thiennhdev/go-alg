package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

// tách duyệt và display khác nhau
// dùng biến leftToRight đánh bool để append display
func zigZagTraversal(root *Node) [][]int {
	res := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		queuSize := len(queue)
		level := make([]int, 0, queuSize)
		for i := 0; i < queuSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.val)

			if cur.left != nil {
				queue = append(queue, cur.left)
			}
			if cur.right != nil {
				queue = append(queue, cur.right)
			}
		}

		// 2 4 bi dao , chan thi dao
		if len(res)%2 != 0 {
			start := 0
			end := len(level) - 1
			for end > start {
				level[start], level[end] = level[end], level[start]
				end--
				start++
			}
		}
		res = append(res, level)
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

	fmt.Println(zigZagTraversal(root))
}
