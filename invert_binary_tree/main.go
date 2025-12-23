package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

// Đảo trái qua phải, phải qua trái của từng node
func invertBinaryTree(tree *Node) *Node {
	if tree == nil {
		return nil
	}
	return &Node{
		val:   tree.val,
		left:  invertBinaryTree(tree.right),
		right: invertBinaryTree(tree.left),
	}
}

func main() {

}
