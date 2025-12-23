package main

import (
	"strconv"
	"strings"
)

type Node struct {
	val      int
	children []Node
}

func dfs(root Node, path []int, res *[]string) {
	path = append(path, root.val)
	if len(root.children) == 0 {
		strArr := make([]string, len(path))
		for i, v := range path {
			strArr[i] = strconv.Itoa(v)
		}

		result := strings.Join(strArr, "->")
		*res = append(*res, result)
		return
	}
	for _, r := range root.children {
		dfs(r, path, res)
	}
}

func ternaryTreePaths(root Node) []string {
	res := make([]string, 0)
	dfs(root, []int{}, &res)
	return res
}

func main() {

	root := Node{
		val: 1,
		children: []Node{
			{
				val: 2,
				children: []Node{
					{
						val: 3,
					},
				},
			},
			{
				val: 4,
			},
			{
				val: 6,
			},
		},
	}

	ternaryTreePaths(root)
}
