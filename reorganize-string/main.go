package main

import (
	"container/heap"
	"fmt"
)

// Example 1:

// Input: s = "aab"
// Output: "aba"
// Example 2:

// Input: s = "aaab"
// Output: ""
type Item struct {
	freq   int
	letter string
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq } // max heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func reorganizeString(s string) string {
	res := ""
	mapFreq := make(map[string]int, 0)
	for _, char := range s {
		mapFreq[fmt.Sprintf("%c", char)] = mapFreq[fmt.Sprintf("%c", char)] + 1
	}

	h := &MaxHeap{}
	heap.Init(h)

	maxFreq := -1
	for letter, freq := range mapFreq {
		if maxFreq < freq {
			maxFreq = freq
		}
		heap.Push(h, Item{
			freq:   freq,
			letter: letter,
		})
	}

	if maxFreq > ((len(s) + 1) / 2) {
		return res
	}

	for h.Len() > 0 {
		max1 := heap.Pop(h).(Item)
		max2 := Item{}
		res = res + max1.letter
		has2 := false
		if h.Len() > 0 {
			max2 = heap.Pop(h).(Item)
			res = res + max2.letter
			has2 = true
		}

		if max1.freq > 1 {
			heap.Push(h, Item{
				letter: max1.letter,
				freq:   max1.freq - 1,
			})
		}

		if has2 && max2.freq > 1 {
			heap.Push(h, Item{
				letter: max2.letter,
				freq:   max2.freq - 1,
			})
		}
	}
	return res
}

func main() {
	fmt.Println(reorganizeString(
		"abbabbaaab",
	))
}

// expect ababababab

// wrong abbaabbaab
