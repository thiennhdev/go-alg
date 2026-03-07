package main

import "container/heap"

// MaxHeap for left half
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap for right half
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor() MedianFinder {
	l := &MaxHeap{}
	r := &MinHeap{}
	heap.Init(l)
	heap.Init(r)
	return MedianFinder{left: l, right: r}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 || num <= (*this.left)[0] {
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	// rebalance: left can have at most 1 more element than right
	if this.left.Len() > this.right.Len()+1 {
		x := heap.Pop(this.left).(int)
		heap.Push(this.right, x)
	} else if this.right.Len() > this.left.Len() {
		x := heap.Pop(this.right).(int)
		heap.Push(this.left, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64((*this.left)[0]) // peek
	}
	// even
	return (float64((*this.left)[0]) + float64((*this.right)[0])) / 2.0
}

func main() {}
