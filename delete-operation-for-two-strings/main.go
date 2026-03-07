package main

import "fmt"

func min(i, j int) int {
	if j < i {
		return j
	}
	return i
}
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	maxLen := len1
	if maxLen < len2 {
		maxLen = len2
	}
	maxLen++
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

		res := 0
		if word1[i-1] == word2[j-1] {
			res = dp(i-1, j-1)
		} else {
			res = min(dp(i-1, j)+1, dp(i, j-1)+1)
		}

		memo[i][j] = res
		return res
	}

	return dp(len1, len2)
}

func deleteString(costs []int, s1 string, s2 string) int {
	len1 := len(s1)
	len2 := len(s2)

	maxLen := len1
	if maxLen < len2 {
		maxLen = len2
	}
	maxLen++
	memo := make([][]int, maxLen)
	for i := 0; i < maxLen; i++ {
		memo[i] = make([]int, maxLen)
		for j := 0; j < maxLen; j++ {
			memo[i][j] = -1
		}
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == 0 && j == 0 {
			return 0
		}
		if i == 0 {
			return dp(0, j-1) + costs[r2[j-1]-'a']
		}

		if j == 0 {
			return dp(i-1, 0) + costs[r1[i-1]-'a']
		}

		if j == 0 {
			return i
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if r1[i-1] == r2[j-1] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			c1 := costs[r1[i-1]-'a']
			c2 := costs[r2[j-1]-'a']
			memo[i][j] = min(dp(i-1, j)+c1, dp(i, j-1)+c2)
		}
		return memo[i][j]
	}

	return dp(len1, len2)
}

func main() {
	// fmt.Println(minDistance("sea", "eat"))
	fmt.Println(
		deleteString([]int{
			1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
			"abb", "bba"),
	)
}
