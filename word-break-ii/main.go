package main

import (
	"math"
	"strings"
)

// Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences in any order.

// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:

// Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// Output: ["cats and dog","cat sand dog"]
// Example 2:

// Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
// Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
// Explanation: Note that you are allowed to reuse a dictionary word.
// Example 3:

// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: []

func wordBreak(s string, wordDict []string) []string {
	result := make([]string, 0)
	min := math.MaxInt
	max := math.MinInt
	mapWordDict := make(map[string]bool)
	for _, word := range wordDict {
		mapWordDict[word] = true
		if len(word) < min {
			min = len(word)
		}

		if len(word) > max {
			max = len(word)
		}
	}

	var dfs func(startIndex int, path []string)
	dfs = func(startIndex int, path []string) {

		if startIndex >= len(s) {
			result = append(result, strings.Join(path, " "))
			return
		}

		for i := startIndex + 1; i <= len(s); i++ {
			temp := s[startIndex:i]
			if len(temp) < min {
				continue
			}

			if len(temp) > max {
				return
			}

			if mapWordDict[temp] {
				dfs(i, append(path, temp))
			}

		}
	}

	dfs(0, []string{})

	return result
}

func main() {

	wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
}
