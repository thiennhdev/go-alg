package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

// root pop root
// left right

func serialize(root *Node) string {
	if root == nil {
		return ""
	}
	result := make([]string, 0)
	stack := make([]*Node, 0)
	// Init first stack
	stack = append(stack, root)
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == nil {
			result = append(result, "x")
		} else {
			result = append(result, strconv.Itoa(current.val))
			stack = append(stack, current.right)
			stack = append(stack, current.left)
		}

	}
	return strings.Join(result, " ")
}

var idx int

func deserializeRec(listNode []string) *Node {
	if idx >= len(listNode) {
		return nil
	}
	if listNode[idx] == "x" {
		idx++
		return nil
	}
	num, _ := strconv.Atoi((listNode[idx]))

	root := &Node{
		val: num,
	}
	idx++
	root.left = deserializeRec(listNode)
	root.right = deserializeRec(listNode)

	return root
}

func deserialize(root string) *Node {
	fmt.Println("root: ", root)
	if len(root) == 0 {
		return nil
	}
	arr := strings.Fields(root)
	idx = 0
	fmt.Println("arr:", arr)
	return deserializeRec(arr)
}

func main() {
	root := &Node{val: 6}
	root.left = &Node{val: 4}
	root.left.left = &Node{val: 3}
	root.left.right = &Node{val: 5}

	root.right = &Node{val: 8}

	s := serialize(root)
	fmt.Println(s)

	fmt.Println(deserialize(s))
}
