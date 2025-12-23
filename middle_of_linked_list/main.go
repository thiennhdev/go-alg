package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func middleOfLinkedList(head *Node) int {
	if head == nil {
		return -1
	}

	if head.next == nil {
		return head.val
	}
	current := head
	fastNode := head
	for current.next != nil && fastNode.next != nil && fastNode.next.next != nil {
		current = current.next
		fastNode = fastNode.next.next
	}

	if fastNode.next != nil {
		current = current.next
	}

	return current.val
}

func main() {
	node0 := &Node{val: 0, next: &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3, next: &Node{val: 4}}}}}
	fmt.Println(middleOfLinkedList(node0))
}
