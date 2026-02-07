package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Pop() chỉ là remove-last
// heap.Pop() mới là remove-min

func mergeKLists(lists []*ListNode) *ListNode {
	h := &minHeap{}
	heap.Init(h)

	for _, node := range lists {
		heap.Push(h, node)
	}
	res := &ListNode{}
	tail := res
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode) // heap pop lấy min
		tail.Next = node
		tail = tail.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	tail.Next = nil
	return res.Next
}

func main() {

	// L1: 1->4->5 (tự nối node)
	// n1 := &ListNode{Val: 1}
	// n2 := &ListNode{Val: 4}
	// n3 := &ListNode{Val: 5}
	// n1.Next = n2
	// n2.Next = n3

	// // L2: 1->3->4
	// n4 := &ListNode{Val: 1}
	// n5 := &ListNode{Val: 3}
	// n6 := &ListNode{Val: 4}
	// n4.Next = n5
	// n5.Next = n6

	// // L3: 2->6
	// n7 := &ListNode{Val: 2}
	// n8 := &ListNode{Val: 6}
	// n7.Next = n8

	n1 := &ListNode{Val: 1}
	n0 := &ListNode{Val: 0}

	lists := []*ListNode{n1, n0}

	merged := mergeKLists(lists)
	fmt.Println(merged) // expected: 1->1->2->3->4->4->5->6
}
