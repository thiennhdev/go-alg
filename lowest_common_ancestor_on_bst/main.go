package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

func lcaOnBst(bst *Node, p int, q int) int {
	if bst == nil {
		return -1
	}

	// left
	if bst.val > p && bst.val > q {
		return lcaOnBst(bst.left, p, q)
	}

	if bst.val < p && bst.val < q {
		return lcaOnBst(bst.right, p, q)
	}

	// split
	if (q < bst.val && p > bst.val) || (q > bst.val && p < bst.val) || (bst.val == q) || bst.val == p {
		return bst.val
	}

	return -1
}

func main() {

}
