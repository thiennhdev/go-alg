package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': true,
		'i': true,
		'e': true,
		'o': true,
		'u': true,
		'A': true,
		'I': true,
		'E': true,
		'O': true,
		'U': true,
	}

	endPtr := len(s) - 1
	startPtr := 0
	chars := []rune(s)
	for endPtr > startPtr {
		for startPtr < endPtr && !vowels[chars[startPtr]] {
			startPtr++
		}

		for endPtr > startPtr && !vowels[chars[endPtr]] {
			endPtr--
		}

		if startPtr >= endPtr {
			break
		}
		chars[startPtr], chars[endPtr] = chars[endPtr], chars[startPtr]
		startPtr++
		endPtr--

	}
	return string(chars)
}
