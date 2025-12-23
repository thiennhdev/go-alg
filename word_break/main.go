package main

import "fmt"

// target = "algomonster"
// words = ["algo", "monster"]
// Output: true

// aab
// a aa b

// for words find start
func wordBreak(target string, words []string) bool {
	// create a map for quick lookup of words
	memoization := make(map[int]bool)

	var dfs func(startIndex int) bool
	dfs = func(startIndex int) bool {
		if startIndex == len(target) {
			return true
		}

		if v, ok := memoization[startIndex]; ok {
			return v
		}
		var ans bool
		for _, word := range words {
			if startIndex+len(word) <= len(target) && target[startIndex:startIndex+len(word)] == word {
				ans = ans || dfs(startIndex+len(word))
				memoization[startIndex+len(word)] = ans
			}
		}
		return ans
	}

	return dfs(0)
}

func main() {
	fmt.Println(wordBreak("algomonster", []string{"algo", "monster"}))
}
