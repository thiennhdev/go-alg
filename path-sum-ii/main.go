package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// smaller -> continue
// bigger return cat path
// equal append
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}

	var dfs func(current *TreeNode, path []int, total int)
	dfs = func(current *TreeNode, path []int, total int) {
		if current == nil {
			return
		}

		if current.Left == nil && current.Right == nil {
			if total+current.Val == targetSum {
				pathCopy := make([]int, len(path))
				copy(pathCopy, path)
				result = append(result, append(pathCopy, current.Val))
			}
			return
		}

		dfs(current.Left, append(path, current.Val), total+current.Val)
		dfs(current.Right, append(path, current.Val), total+current.Val)
	}

	dfs(root, []int{}, 0)
	return result
}

func createTree() *TreeNode {
	// Root: 5
	root := &TreeNode{Val: 5}

	// Level 1: 4 (left), 8 (right)
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}

	// Level 2:
	// - 4.Left = 11
	root.Left.Left = &TreeNode{Val: 11}

	// - 8.Left = 13, 8.Right = 4
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}

	// Level 3:
	// - 11.Left = 2
	root.Left.Left.Left = &TreeNode{Val: 2}

	// - 4 (right branch).Left = 5, 4 (right branch).Right = 1
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}

	return root
}

func main() {
	root := createTree()

	// Test với pathSum
	fmt.Println("\npathSum với targetSum = 22:", pathSum(root, 22))
}
