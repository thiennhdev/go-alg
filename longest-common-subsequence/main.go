package main

import "fmt"

// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.

func max(i, j int) int {
	if j > i {
		return j
	}
	return i
}
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)

	maxLen := max(len1, len2) + 1

	// traversal form last
	memo := make([][]int, maxLen)
	for i := 0; i < maxLen; i++ {
		memo[i] = make([]int, maxLen)
		for j := 0; j < maxLen; j++ {
			memo[i][j] = -1
		}
	}

	var lcs func(i, j int) int
	lcs = func(i, j int) int {
		if i == 0 || j == 0 {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := 0
		if text1[i-1] == text2[j-1] {
			res = 1 + lcs(i-1, j-1)
		} else {
			res = max(lcs(i-1, j), lcs(i, j-1))
		}

		memo[i][j] = res
		return res
	}

	return lcs(len1, len2)
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
