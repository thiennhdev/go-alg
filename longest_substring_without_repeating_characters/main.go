package main

import "fmt"

func longestSubstringWithoutRepeatingCharacters(s string) int {
	count := 0
	mapCh := map[rune]int{}
	for i, ch := range s {
		if _, ok := mapCh[ch]; !ok {
			mapCh[ch] = i
			if len(mapCh) > count {
				count = len(mapCh)
			}
		} else {
			index := mapCh[ch] + 1
			mapCh = map[rune]int{}
			for index <= i {
				mapCh[rune(s[index])] = index
				index++
			}
		}
	}
	return count
}

// add tu i den last index bi trung
func main() {
	fmt.Println(longestSubstringWithoutRepeatingCharacters("abcdbea"))
}

// a b c d [b] => co b day ptr trung + 1
// map character to last index
