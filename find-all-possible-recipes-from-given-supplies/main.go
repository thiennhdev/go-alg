package main

import "fmt"

// Example 1:

// Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
// Output: ["bread"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// Example 2:

// Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
// Output: ["bread","sandwich"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// ingredient -> list of recipes that depend on it
	canCreate := make(map[string][]string, len(recipes))
	indegree := make(map[string]int, len(recipes))

	for i, recipe := range recipes {
		indegree[recipe] = len(ingredients[i])
		for _, ing := range ingredients[i] {
			canCreate[ing] = append(canCreate[ing], recipe)
		}
	}

	res := make([]string, 0, len(recipes))
	queue := make([]string, 0, len(supplies)+len(recipes))
	queue = append(queue, supplies...)

	for head := 0; head < len(queue); head++ { // dùng head index để khỏi pop O(n)
		cur := queue[head]

		for _, r := range canCreate[cur] {
			indegree[r]--
			if indegree[r] == 0 {
				res = append(res, r)
				queue = append(queue, r) // recipe thành supply mới
			}
		}
	}

	return res
}

func main() {
	recipes := []string{"bread", "sandwich"}

	ingredients := [][]string{
		{"yeast", "flour"}, // cho bread
		{"bread", "meat"},  // cho sandwich
	}

	supplies := []string{"yeast", "flour", "meat"}

	fmt.Println(findAllRecipes(recipes, ingredients, supplies))
	// Output: [bread sandwich]
}
