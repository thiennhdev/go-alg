package main

import (
	"fmt"
	"strings"
)

func generateParentheses(n int) []string {
	var res []string
	var backtracking func(int, int, int, []string, *[]string)

	backtracking = func(n int, openCount, closedCount int, path []string, res *[]string) {
		if len(path) == 2*n {
			*res = append(*res, strings.Join(path, ""))
			return
		}

		// Callback
		if openCount < n {
			backtracking(n, openCount+1, closedCount, append(path, "("), res)
		}
		if closedCount < openCount {
			backtracking(n, openCount, closedCount+1, append(path, ")"), res)
		}
	}

	backtracking(n, 0, 0, []string{}, &res)
	return res
}

func main() {
	fmt.Println("generateParentheses : ", generateParentheses(3))
}
