package main

import "fmt"

// Dont suggest to use this code, it is just for learning purposes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

type BST struct {
	root *TreeNode
}

func NewBST() *BST {
	return &BST{}
}

// Dont suggest to use this code, it is just for learning purposes
func (b *BST) Insert(val int) {
	if b.root == nil {
		b.root = NewTreeNode(val)
		return
	}
	b.insertNode(b.root, val)
}

func (b *BST) insertNode(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = NewTreeNode(val)
			return
		}
		b.insertNode(root.Left, val)
	} else if val > root.Val {
		if root.Right == nil {
			root.Right = NewTreeNode(val)
			return
		}
		b.insertNode(root.Right, val)
	}
}

func (b *BST) Find(val int) bool {
	if b.root == nil {
		return false
	}

	findNode := b.findNode(b.root, val)
	return findNode != nil
}

func (b *BST) findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return b.findNode(root.Left, val)
	}

	if val > root.Val {
		return b.findNode(root.Right, val)
	}

	return nil
}

func (b *BST) findParent(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if (root.Left != nil && root.Left.Val == val) || (root.Right != nil && root.Right.Val == val) {
		return root
	}

	if val < root.Val {
		return b.findParent(root.Left, val)
	}

	if val > root.Val {
		return b.findParent(root.Right, val)
	}

	return nil
}

func (b *BST) findMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return b.findMinNode(root.Left)
}

func (b *BST) Delete(val int) {
	// 1. Find the node to delete
	nodeToDelete := b.findNode(b.root, val)
	if nodeToDelete == nil {
		return
	}
	fmt.Println("nodeToDelete:", nodeToDelete.Val)
	isRoot := b.root.Val == val
	// 2. Find the parent of the node to delete
	parent := b.findParent(b.root, val)
	if parent == nil && !isRoot {
		return
	}

	// 3. Delete the node
	// 3.1. Node has no children
	if !isRoot && nodeToDelete.Left == nil && nodeToDelete.Right == nil {
		if parent.Left.Val == nodeToDelete.Val {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}

	// 3.2. Node has one child
	// Node has left child
	if !isRoot && nodeToDelete.Left != nil && nodeToDelete.Right == nil {
		if parent.Left.Val == nodeToDelete.Val {
			parent.Left = nodeToDelete.Left
		} else {
			parent.Right = nodeToDelete.Left
		}
		return
	}

	// Node has right child
	if !isRoot && nodeToDelete.Left == nil && nodeToDelete.Right != nil {
		if parent.Left.Val == nodeToDelete.Val {
			parent.Left = nodeToDelete.Right
		} else {
			parent.Right = nodeToDelete.Right
		}
		return
	}

	// 3.3. Node has two children
	if nodeToDelete.Left != nil && nodeToDelete.Right != nil {
		successor := b.findMinNode(nodeToDelete.Right)
		if successor == nodeToDelete.Right {
			nodeToDelete.Val = successor.Val
			nodeToDelete.Right = successor.Right
		} else {
			b.Delete(successor.Val)
			nodeToDelete.Val = successor.Val
		}
	}
}

func (b *BST) Print() {
	fmt.Println("=== BST Structure ===")
	b.printNode(b.root, "", "ROOT")
	fmt.Println("====================")
}

func (b *BST) printNode(node *TreeNode, prefix string, position string) {
	if node == nil {
		return
	}

	// In node hiện tại với vị trí rõ ràng
	fmt.Printf("%s%s: %d\n", prefix, position, node.Val)

	// Tạo prefix cho các node con
	childPrefix := prefix + "    "

	// In node con trái trước
	if node.Left != nil {
		b.printNode(node.Left, childPrefix, "L")
	} else if node.Right != nil {
		// Chỉ hiển thị "L: null" nếu có node phải
		fmt.Printf("%sL: null\n", childPrefix)
	}

	// In node con phải sau
	if node.Right != nil {
		b.printNode(node.Right, childPrefix, "R")
	} else if node.Left != nil {
		// Chỉ hiển thị "R: null" nếu có node trái
		fmt.Printf("%sR: null\n", childPrefix)
	}
}
func main() {
	bst := NewBST()
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	// Print the tree
	bst.Print()

	bst.Delete(50)
	bst.Print()

}
