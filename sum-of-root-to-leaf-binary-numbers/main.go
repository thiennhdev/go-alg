package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binarySliceToDecimal(bits []int) int {
	result := 0
	for _, bit := range bits {
		result = result*2 + bit
	}
	return result
}

func sumRootToLeaf(root *TreeNode) int {

	var dfs func(cur *TreeNode, path []int)
	result := 0
	dfs = func(cur *TreeNode, path []int) {
		if cur.Left == nil && cur.Right == nil {
			result += binarySliceToDecimal(append(path, cur.Val))
			return
		}

		if cur.Left != nil {
			dfs(cur.Left, append(path, cur.Val))
		}

		if cur.Right != nil {
			dfs(cur.Right, append(path, cur.Val))
		}
	}

	dfs(root, []int{})

	return result
}

// build tree from level-order array where nil is represented by pointer nil in []*int
func buildTreeLevelOrder(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	q := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(q) > 0 {
		cur := q[0]
		q = q[1:]

		// left
		if i < len(arr) && arr[i] != nil {
			cur.Left = &TreeNode{Val: *arr[i]}
			q = append(q, cur.Left)
		}
		i++

		// right
		if i < len(arr) && arr[i] != nil {
			cur.Right = &TreeNode{Val: *arr[i]}
			q = append(q, cur.Right)
		}
		i++
	}

	return root
}
func main() {
	// Input: root = [1,0,1,0,1,0,1]
	v1 := 1
	v0 := 0
	arr := []*int{&v1, &v0, &v1, &v0, &v1, &v0, &v1}

	root := buildTreeLevelOrder(arr)
	fmt.Println(sumRootToLeaf(root)) // expected: 22
}
