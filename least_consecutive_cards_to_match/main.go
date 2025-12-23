package main

import "fmt"

func leastConsecutiveCardsToMatch(cards []int) int {
	fast := 0
	count := -1
	mapCard := map[int]int{} // card => index
	for fast < len(cards) {
		cardFast := cards[fast]
		if _, ok := mapCard[cardFast]; !ok {
			mapCard[cardFast] = fast
		} else {
			if count < len(mapCard)+1 {
				count = len(mapCard) + 1
			}
			index := mapCard[cardFast] + 1
			mapCard = map[int]int{}
			for index <= fast {
				cardTemp := cards[index]
				mapCard[cardTemp] = index
				index++
			}
		}
		fast++
	}

	return count
}

func main() {
	fmt.Println(leastConsecutiveCardsToMatch(([]int{3, 4, 2, 3, 4, 7})))
}
