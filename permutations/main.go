package main

import "fmt"

// abc
// a
// ab ac
// abc

// b
// ba bc
// bac bca

// create new string by append path

// func newRest(letters, s string) string {
// 	index := -1
// 	for i, v := range letters {
// 		if v ==
// 	}
// }

// abc
func newRest(letters, cur string) string {
	used := make(map[rune]int)
	for _, c := range cur {
		used[c]++
	}

	var result []rune
	for _, c := range letters {
		if used[c] > 0 {
			used[c]--
		} else {
			result = append(result, c)
		}
	}

	return string(result)
}

func permutations(letters string) []string {
	var res []string
	var dfs func(string, string, []string)
	dfs = func(letters string, cur string, path []string) {
		if len(letters) == len(path) {
			res = append(res, cur)
			return
		}

		newRestLetter := newRest(letters, cur)
		for _, ch := range newRestLetter {
			newCur := cur + string(ch)
			dfs(letters, newCur, append(path, newCur))
		}
	}

	dfs(letters, "", []string{})
	return res
}

func main() {
	fmt.Println("permutations: ", permutations("abc"))
}
