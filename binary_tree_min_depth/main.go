package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

// find depth
// leaf dont have left or right child
func binaryTreeMinDepth(root *Node) int {
	depth := 0
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.left == nil && cur.right == nil {
				return depth
			}

			if cur.left != nil {
				queue = append(queue, cur.left)
			}

			if cur.right != nil {
				queue = append(queue, cur.right)
			}
		}

		depth++
	}
	return depth
}
func main() {

}
