package main

type Node struct {
	val  int
	next *Node
}

func hasCycle(nodes *Node) bool {
	fast := nodes
	slow := nodes

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast != nil && slow != nil && fast == slow {
			return true
		}
	}
	return false
}

func main() {

}
