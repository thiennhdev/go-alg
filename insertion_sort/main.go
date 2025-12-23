package main

import "fmt"

func insertionSort(arr []int) {
	for i, _ := range arr {
		currentIndex := i
		for currentIndex > 0 && arr[currentIndex] < arr[currentIndex-1] {
			arr[currentIndex], arr[currentIndex-1] = arr[currentIndex-1], arr[currentIndex]
			currentIndex--
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	i := 0
	for i < n {
		swapped := false
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
		i++
	}
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	bubbleSort(arr)
	fmt.Println(arr)
}
