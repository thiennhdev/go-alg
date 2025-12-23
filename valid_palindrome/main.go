package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	startPtr, endPtr := 0, len(s)-1
	for endPtr >= startPtr {
		for !unicode.IsLetter(rune(s[startPtr])) {
			startPtr++
		}

		for !unicode.IsLetter(rune(s[endPtr])) {
			endPtr--
		}

		if unicode.ToLower(rune(s[startPtr])) != unicode.ToLower(rune(s[endPtr])) {
			return false
		}
		startPtr++
		endPtr--
	}
	return true
}

func main() {
	fmt.Println("isPalindrome: ", isPalindrome("Do geese see God?"))
}
