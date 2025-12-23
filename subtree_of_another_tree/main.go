package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

func isSameTree(root *Node, subRoot *Node) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	return root.val == subRoot.val && isSameTree(root.left, subRoot.left) && isSameTree(root.right, subRoot.right)
}

func subtreeOfAnotherTree(root *Node, subRoot *Node) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	return isSameTree(root, subRoot) || subtreeOfAnotherTree(root.left, subRoot) || subtreeOfAnotherTree(root.right, subRoot)
}

func main() {

}
