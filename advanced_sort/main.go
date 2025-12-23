package main

import "fmt"

func mergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}
	if n == 0 {
		return []int{}
	}
	mid := n / 2

	a := mergeSort(arr[:mid])
	b := mergeSort(arr[mid:])

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	// Merge
	i := 0 // a
	j := 0 // b
	ptr := 0
	z := len(a) + len(b)
	mergeArr := make([]int, z, z)
	for i < len(a) || j < len(b) {
		if j < len(b) && i < len(a) {
			if a[i] < b[j] {
				mergeArr[ptr] = a[i]
				i++
			} else {
				mergeArr[ptr] = b[j]
				j++
			}
		} else if j < len(b) && i >= len(a) {
			mergeArr[ptr] = b[j]
			j++
		} else if i < len(a) && j >= len(b) {
			mergeArr[ptr] = a[i]
			i++
		}
		ptr++
	}
	return mergeArr
}

// pivot is end
func quickSort(arr []int) []int {
	sortListInterval(arr, 0, len(arr))
	return arr
}

func sortListInterval(unsortedList []int, start, end int) {
	// Must more than 2 elements
	if end-start <= 1 {
		return
	}

	startPtr := start
	endPtr := end - 1

	// Pick pivot is last element
	pivotValue := unsortedList[end-1]
	for endPtr > startPtr {

		// find first element is smaller than pivot
		for endPtr > startPtr && unsortedList[endPtr] >= pivotValue {
			endPtr--
		}

		// find first element is greater than pivot
		for startPtr < endPtr && unsortedList[startPtr] < pivotValue {
			startPtr++
		}

		//swap
		unsortedList[startPtr], unsortedList[endPtr] = unsortedList[endPtr], unsortedList[startPtr]
	}

	// Swap pivot
	unsortedList[end-1], unsortedList[startPtr+1] = unsortedList[startPtr+1], unsortedList[end-1]
	sortListInterval(unsortedList, start, startPtr)
	sortListInterval(unsortedList, startPtr+1, end)
}

func sortListIntervalUnfix(unsortedList []int, start, end int) {
	// If segment is 1 or 0, it's sorted
	if end-start <= 1 {
		return
	}

	// Using last element as the pivot
	pivot := unsortedList[end-1]
	startPtr, endPtr := start, end-1

	// Partitioning process
	for startPtr < endPtr {
		// Find the next element from the left that is larger than the pivot
		for unsortedList[startPtr] < pivot && startPtr < endPtr {
			startPtr++
		}

		// Find the next element from the right that is smaller than or equal to the pivot
		for unsortedList[endPtr] >= pivot && startPtr < endPtr {
			endPtr--
		}

		// Swap if pointers haven't met
		unsortedList[startPtr], unsortedList[endPtr] = unsortedList[endPtr], unsortedList[startPtr]
	}

	// Place pivot in its final position
	unsortedList[startPtr], unsortedList[end-1] = unsortedList[end-1], unsortedList[startPtr]

	// Sort left and right of the pivot
	sortListInterval(unsortedList, start, startPtr)
	sortListInterval(unsortedList, startPtr+1, end)
}

func sortListIntervalOriginal(unsortedList []int, start, end int) {
	// If segment is 1 or 0, it's sorted
	if end-start <= 1 {
		return
	}

	// Using last element as the pivot
	pivot := unsortedList[end-1]
	startPtr, endPtr := start, end-1

	// Partitioning process
	for startPtr < endPtr {
		// Find the next element from the left that is larger than the pivot
		for unsortedList[startPtr] < pivot && startPtr < endPtr {
			startPtr++
		}

		// Find the next element from the right that is smaller than or equal to the pivot
		for unsortedList[endPtr] >= pivot && startPtr < endPtr {
			endPtr--
		}

		// Swap if pointers haven't met
		unsortedList[startPtr], unsortedList[endPtr] = unsortedList[endPtr], unsortedList[startPtr]
	}

	// Place pivot in its final position
	unsortedList[startPtr], unsortedList[end-1] = unsortedList[end-1], unsortedList[startPtr]

	// Sort left and right of the pivot
	sortListInterval(unsortedList, start, startPtr)
	sortListInterval(unsortedList, startPtr+1, end)
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	arr2 := quickSort(arr)
	fmt.Println(arr2)
}
