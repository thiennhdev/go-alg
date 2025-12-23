package main

import (
	"fmt"
	"strings"
)

func dfs(n, max int, path []string, res *[]string) {
	if max == n {
		*res = append(*res, strings.Join(path, ""))
		return
	}

	dfs(n+1, max, append(path, "a"), res)
	dfs(n+1, max, append(path, "b"), res)

}

func letterCombination(n int) []string {
	res := make([]string, 0)
	dfs(0, n, []string{}, &res)
	return res
}

func main() {
	fmt.Println("letterCombination: ", letterCombination(2))
}
