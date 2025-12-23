package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	startIndex := 0
	endIndex := len(s) - 1
	for startIndex < endIndex {
		if s[startIndex] != s[endIndex] {
			return false
		}
		startIndex++
		endIndex--
	}

	return true
}

func backtracking(s string, index int, path []string, res *[][]string) {
	if index == len(s) {
		*res = append(*res, path)
		return
	}

	for end := index + 1; end <= len(s); end++ {
		temp := s[index:end]
		if isPalindrome(temp) {
			backtracking(s, end, append(path, temp), res)
		}
	}
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	backtracking(s, 0, []string{}, &res)
	return res
}

func main() {
	fmt.Println(partition("aab"))
}
