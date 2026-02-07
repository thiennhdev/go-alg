package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func nthUglyNumber(n int) int {
	h := &MinHeap{
		1, // init
	}

	vistied := make(map[int]bool)
	vistied[1] = true
	heap.Init(h)
	for i := 1; i < n; i++ {
		cur := heap.Pop(h).(int)
		mul := []int{
			2, 3, 5,
		}
		for _, m := range mul {
			if vistied[m*cur] {
				continue
			}
			heap.Push(h, m*cur)
			vistied[m*cur] = true
		}

	}

	return heap.Pop(h).(int)
}

func main() {
	fmt.Println("nthUglyNumber: ", nthUglyNumber(10))
}
