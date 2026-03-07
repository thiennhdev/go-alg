package main

import "fmt"

// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')

// from last
// Insert a character
// Delete a character
// Replace a character
// from word 1 -> word 2
func min(i, j int) int {
	if j < i {
		return j
	}
	return i
}

func min3(i, j, x int) int {

	return min(min(i, j), x)
}

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	maxLen := len1
	if len2 > maxLen {
		maxLen = len2
	}
	maxLen = maxLen + 1
	memo := make([][]int, maxLen)
	for i := 0; i < maxLen; i++ {
		memo[i] = make([]int, maxLen)
		for j := 0; j < maxLen; j++ {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int

	dp = func(i, j int) int {
		if i == 0 {
			return j
		}

		if j == 0 {
			return i
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		cost := 1
		if word1[i-1] == word2[j-1] {
			cost = 0
		}

		memo[i][j] = min3(
			dp(i-1, j)+1,      // delete
			dp(i, j-1)+1,      //  insert
			dp(i-1, j-1)+cost, // match replace
		)

		return memo[i][j]
	}

	return dp(len1, len2)

}

func main() {
	fmt.Println(minDistance("intention", "execution"))

}
