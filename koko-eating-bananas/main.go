package main

import (
	"fmt"
	"math"
)

/*
Input: piles = [3,6,7,11], h = 8
Output: 4
*/

// có 4 đống chuối mỗi đống có n quả => cho 8 tiếng, mỗi tiếng chỉ được ăn 1 đống, và mỗi tiếng phải ăn bao nhiêu qủa để xong trong 8 tiếng

func feasible(piles []int, limitTime int, canEatNumBanana int) bool {
	totalTime := 0
	for _, p := range piles {
		timeEat := math.Ceil(float64(p) / float64(canEatNumBanana))
		totalTime += int(timeEat)
	}
	return totalTime <= limitTime
}

func minEatingSpeed(piles []int, h int) int {
	bound := -1

	// simulator time
	startPtr, endPtr := 1, 60

	for endPtr >= startPtr {
		mid := (startPtr + endPtr) / 2
		if feasible(piles, h, mid) { // có thể ăn => giảm time xuống
			bound = mid
			endPtr = mid - 1
		} else {
			startPtr = mid + 1
		}
	}

	return bound
}

func main() {
	fmt.Println("minEatingSpeed: ", minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}
