package main

import "fmt"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

type AVL struct {
	root *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:    val,
		Height: 0,
	}
}

func NewAVL(val int) *AVL {
	return &AVL{
		root: NewNode(val),
	}
}

func (a *AVL) Insert(value int) {
	a.root = a.insertNode(a.root, value)
}

func (a *AVL) insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return NewNode(value)
	}

	if value < node.Val {
		node.Left = a.insertNode(node.Left, value)
	} else if value > node.Val {
		node.Right = a.insertNode(node.Right, value)
	} else {
		return node
	}

	node.Height = 1 + max(a.getHeight(node.Left), a.getHeight(node.Right))
	balance := a.getBalanceFactor(node)

	// Left Left Case
	if balance > 1 && value < node.Left.Val {
		return a.rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Val {
		return a.rotateLeft(node)
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Val {
		node.Left = a.rotateLeft(node.Left)
		return a.rotateRight(node)
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Val {
		node.Right = a.rotateRight(node.Right)
		return a.rotateLeft(node)
	}
	return node
}

func (a *AVL) getHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(a.getHeight(root.Left), a.getHeight(root.Right))
}

// left - right
//
//	if positive, left heavy
func (a *AVL) getBalanceFactor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return a.getHeight(root.Left) - a.getHeight(root.Right)
}

func (a *AVL) rotateRight(root *TreeNode) *TreeNode {
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root
	root.Height = 1 + max(a.getHeight(root.Left), a.getHeight(root.Right))
	newRoot.Height = 1 + max(a.getHeight(newRoot.Left), a.getHeight(newRoot.Right))
	return newRoot
}

func (a *AVL) rotateLeft(root *TreeNode) *TreeNode {
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root
	root.Height = 1 + max(a.getHeight(root.Left), a.getHeight(root.Right))
	newRoot.Height = 1 + max(a.getHeight(newRoot.Left), a.getHeight(newRoot.Right))
	return newRoot
}

//  x
// / \
// y   z
// /
// z

func main() {
	avl := NewAVL(10)
	avl.Insert(20)
	avl.Insert(30)
	avl.Insert(40)
	avl.Insert(50)
	avl.Insert(35)

	fmt.Println("done: ", avl.root)
}
