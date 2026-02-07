package main

import "container/heap"

// Input: nums = [3,2,1,5,6,4], k = 2
// Output: 5

// min heap
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
			continue
		}

		min := (*h)[0]
		if min < num {
			heap.Pop(h)
			heap.Push(h, num)
		}

	}

	return heap.Pop(h).(int)
}

func main() {

}
