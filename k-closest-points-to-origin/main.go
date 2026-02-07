package main

import (
	"container/heap"
	"fmt"
)

type item struct {
	dist int
	p    []int
}

// max-heap theo dist (dist lớn hơn nằm trên)
type maxHeap []item

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].dist > h[j].dist } // max-heap
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func dist2(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}

func kClosest(points [][]int, k int) [][]int {
	if k <= 0 {
		return [][]int{}
	}
	if k >= len(points) {
		return points
	}

	h := &maxHeap{}
	heap.Init(h)

	for _, p := range points {
		d := dist2(p)
		if h.Len() < k {
			heap.Push(h, item{dist: d, p: p})
			continue
		}

		if d < (*h)[0].dist {
			heap.Pop(h)
			heap.Push(h, item{dist: d, p: p})
		}
	}

	res := make([][]int, h.Len())
	for i := 0; h.Len() > 0; i++ {
		res[i] = heap.Pop(h).(item).p
	}
	return res
}

func main() {
	fmt.Println(kClosest(
		[][]int{
			{1, 1},
			{2, 2},
			{3, 3},
		}, 1,
	))
}
