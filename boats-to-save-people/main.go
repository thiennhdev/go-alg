package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{2, 2, 2, 6, 6, 7, 10, 10, 10, 11, 12, 13, 18, 22, 26, 33, 41, 47, 49, 50}, 50))
}

// [1,2] | 3
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	numberOfBouts := 0

	start := 0
	end := len(people) - 1

	for start < end {
		if people[start]+people[end] <= limit {
			numberOfBouts++
			start++
		} else {
			numberOfBouts++
		}
		end--
	}

	if start == end {
		numberOfBouts++
	}

	return numberOfBouts
}
