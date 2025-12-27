package main

import "fmt"

func numSteps(targetCombo string, trappedCombos []string) int {
	nextDigit := make(map[int]int)
	prevDigit := make(map[int]int)
	for i := 0; i <= 9; i++ {
		increase := i + 1
		if increase > 9 {
			increase = 0
		}
		nextDigit[i] = increase

		descrese := i - 1
		if descrese < 0 {
			descrese = 9
		}
		prevDigit[i] = descrese
	}

	mapTrapped := make(map[string]bool)
	for _, trap := range trappedCombos {
		mapTrapped[trap] = true
	}
	queue := make([]string, 0)
	queue = append(queue, "0000")
	steps := map[string]int{
		"0000": 0,
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if mapTrapped[curr] {
			return -1
		}

		if curr == targetCombo {
			return steps[curr]
		}
		for i := 0; i < 4; i++ {
			digit := int(curr[i] - '0')
			nextCombo := fmt.Sprintf("%s%d%s", curr[0:i], nextDigit[digit], curr[i+1:])
			if _, ok := steps[nextCombo]; !ok && !mapTrapped[nextCombo] {
				steps[nextCombo] = steps[curr] + 1
				if nextCombo == targetCombo {
					return steps[nextCombo]
				}
				queue = append(queue, nextCombo)
			}

			prevCombo := fmt.Sprintf("%s%d%s", curr[0:i], prevDigit[digit], curr[i+1:])
			if _, ok := steps[prevCombo]; !ok && !mapTrapped[prevCombo] {
				steps[prevCombo] = steps[curr] + 1
				if prevCombo == targetCombo {
					return steps[prevCombo]
				}
				queue = append(queue, prevCombo)
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(numSteps("0202", []string{"0201", "0101", "0102", "1212", "2002"}))
}
