package main

import "fmt"

func equals(original string, check string) bool {
	mapOriginal := make(map[rune]int)
	for _, o := range original {
		if _, ok := mapOriginal[o]; !ok {
			mapOriginal[o] = 1
		} else {
			mapOriginal[o]++
		}
	}

	for _, c := range check {
		if _, ok := mapOriginal[c]; !ok {
			return false
		} else {
			if mapOriginal[c] == 0 {
				return false
			}
			mapOriginal[c]--
		}
	}
	return true
}

func findAllAnagrams(original string, check string) []int {
	slideCount := len(check) - 1
	result := make([]int, 0)
	for i := slideCount; i < len(original); i++ { // a[b]
		slide := string(original[i-slideCount : i+1])
		if equals(slide, check) {
			result = append(result, i-slideCount)
		}

	}
	return result
}

func main() {
	fmt.Println(findAllAnagrams("cbaebabacd", "abc"))
}
