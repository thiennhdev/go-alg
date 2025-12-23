package main

import (
	"fmt"
)

func main() {
	fmt.Println(newspapersSplit([]int{1, 4, 4}, 3))
}

func feasible(timeLimitPerOne int, paperReadTimes []int, numCoworkers int) bool {
	usedCoworker := 1
	count := 0
	for _, paperReadTime := range paperReadTimes {
		if paperReadTime > timeLimitPerOne {
			return false
		}
		count += paperReadTime
		if count > timeLimitPerOne {
			usedCoworker++
			count = paperReadTime
		}
	}

	return numCoworkers >= usedCoworker
}

func newspapersSplit(newspapersReadTimes []int, numCoworkers int) int {
	minTime := 0
	maxTime := 10 //math.MaxInt32
	bound := -1
	for minTime <= maxTime {
		timeRead := (maxTime + minTime) / 2
		if feasible(timeRead, newspapersReadTimes, numCoworkers) {
			bound = timeRead
			maxTime = timeRead - 1
		} else {
			minTime = timeRead + 1
		}
	}

	return bound
}
