package main

import "fmt"

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func build_tree(s string) *TreeNode {
	skipValue := ' '
	nilNode := 'x'
	head := &TreeNode{}
	previous := make([]*TreeNode, 0)
	current := head
	for _, ch := range s {
		if ch == skipValue {
			fmt.Println("empty: ", string(ch))
			continue
		}

		if ch == nilNode {
			current = nil
			temp := previous[len(previous)-1]
			previous = previous[0 : len(previous)-1]
			temp.Right = &TreeNode{}
			current = temp.Right
			continue
		}
		current.Val = string(ch)
		previous = append(previous, current)
		current.Left = &TreeNode{}
		current = current.Left
	}

	return head
}

// previous : 5 4 3

// pre root-left-right
func main() {
	build_tree("5 4 3 x x 8 x x 6 x x")
}
