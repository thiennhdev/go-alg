package main

import "fmt"

func sortedSquares(nums []int) []int {
	minIndex := 0
	maxIndex := len(nums) - 1
	newNums := make([]int, len(nums))
	i := len(nums) - 1
	for maxIndex >= minIndex {
		minSquare := nums[minIndex] * nums[minIndex]
		maxSquare := nums[maxIndex] * nums[maxIndex]
		if maxSquare > minSquare {
			newNums[i] = maxSquare
			i = i - 1
			maxIndex = maxIndex - 1
		} else {
			newNums[i] = minSquare
			i = i - 1
			minIndex = minIndex + 1
		}
	}
	return newNums
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
