package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("abca"))
}

// eedede
// e eded e
// e e de d e => đẩy 2 đầu
func isPalindrome(startPtr, endPtr int, chars []rune) bool {
	if startPtr > endPtr {
		return false
	}
	for endPtr > startPtr {
		if chars[endPtr] != chars[startPtr] {
			return false
		}
		startPtr++
		endPtr--
	}

	return true
}

func validPalindrome(s string) bool {
	chars := []rune(s)
	endPtr := len(chars) - 1
	for i := 0; i < len(chars); i++ {
		if chars[i] != chars[endPtr] {
			checkOne := isPalindrome(i+1, endPtr, chars)
			checkTwo := isPalindrome(i, endPtr-1, chars)
			return checkOne || checkTwo
		}
		endPtr--
	}

	return true
}
