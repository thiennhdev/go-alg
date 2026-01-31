package main

// Input: equations = ["a==b","b!=a"]
// Output: false

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			parent[y] = x
		}
	}

	for _, equation := range equations {
		if equation[1:3] == "==" {
			x, y := int(equation[0]-'a'), int(equation[3]-'a')
			union(x, y)
		}
	}

	for _, equation := range equations {
		if equation[1:3] == "!=" {
			x, y := int(equation[0]-'a'), int(equation[3]-'a')
			if find(x) != find(y) {
				return false
			}
		}
	}
	return true
}

func main() {

}
