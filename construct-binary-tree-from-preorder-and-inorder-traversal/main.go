package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7].
// root left right ||| left root rigth
//
//	|| càng gần 3 càng gần root
func constructBinaryTree(preorder, inorder []int) *Node {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// root là phần tử đầu của preorder
	rootVal := preorder[0]

	// tìm root trong inorder
	idx := -1
	for i, v := range inorder {
		if v == rootVal {
			idx = i
			break // QUAN TRỌNG: dừng ngay khi tìm thấy
		}
	}
	if idx == -1 { // guard nếu input không hợp lệ
		return nil
	}

	leftSize := idx
	root := &Node{val: rootVal}

	// inorder: [0:idx] là trái, [idx+1:] là phải
	// preorder: [1:1+leftSize] trái, [1+leftSize:] phải
	root.left = constructBinaryTree(preorder[1:1+leftSize], inorder[:idx])
	root.right = constructBinaryTree(preorder[1+leftSize:], inorder[idx+1:])

	return root
}
func main() {

	constructBinaryTree([]int{7, -10, -4, 3, -1, 2, -8, 11}, []int{-4, -10, 3, -1, 7, 11, -8, 2})
}
