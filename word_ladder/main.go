package main

import "fmt"

func compare(root, dst string) bool {
	if len(root) != len(dst) {
		return false
	}
	countWrong := 0
	for i := 0; i < len(root); i++ {
		if root[i] != dst[i] {
			countWrong++
		}

		if countWrong >= 2 {
			return false
		}
	}
	return true
}

func wordLadder(begin string, end string, wordList []string) int {
	queue := make([]string, 0)
	queue = append(queue, begin)
	step := map[string]int{
		begin: 1,
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, word := range wordList {
			if _, ok := step[word]; !ok && compare(cur, word) {
				step[word] = step[cur] + 1
				if word == end {
					return step[word]
				}
				queue = append(queue, word)
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(wordLadder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
