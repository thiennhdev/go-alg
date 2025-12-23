package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

func insertBst(bst *Node, val int) *Node {
	if bst == nil {
		return &Node{
			val: val,
		}
	}

	if val > bst.val {
		bst.right = insertBst(bst.right, val)
	}

	if val < bst.val {
		bst.left = insertBst(bst.left, val)
	}

	return bst
}

func main() {

}
