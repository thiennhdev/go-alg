package main

import "container/heap"

type Node struct {
	val int
	r   int
	c   int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Node)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])

	h := &MinHeap{}
	heap.Init(h)

	for r := 0; r < n; r++ {
		heap.Push(h, Node{val: matrix[r][0], r: r, c: 0})
	}

	// pop k-1 times
	for i := 0; i < k-1; i++ {
		cur := heap.Pop(h).(Node)
		if cur.c+1 < m {
			heap.Push(h, Node{val: matrix[cur.r][cur.c+1], r: cur.r, c: cur.c + 1})
		}
	}

	return heap.Pop(h).(Node).val
}
