package main

import "fmt"

func subarraySum(arr []int, target int) []int {
	mapFindSum := map[int]int{
		0: 0,
	} // map value to index
	beforeValue := 0
	for i := 0; i < len(arr); i++ {
		beforeValue += arr[i]
		check := beforeValue - target
		if _, ok := mapFindSum[check]; !ok {
			mapFindSum[beforeValue] = i + 1
		} else {
			return []int{mapFindSum[check], i + 1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println("subarraySum: ", subarraySum([]int{1, 1, 1}, 2))
}
