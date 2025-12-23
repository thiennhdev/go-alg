package main

import (
	"fmt"
	"strconv"
)

const GapIndex = 64

func decodeWays(digits string) int {
	min := rune('A')
	max := rune('Z')

	var dfs func(current string, rest string)
	var result int
	dfs = func(current, rest string) {
		if rest == "" {
			result++
			return
		}
		// temp := current
		for i := 0; i < len(rest); i++ {
			sub := rest[:i+1]
			n, _ := strconv.Atoi(sub)
			if n == 0 {
				return
			}
			n += GapIndex
			if n >= int(min) && n <= int(max) {
				dfs(current+string(rune(n)), rest[i+1:])
			}
		}
	}
	dfs("", digits)
	return result
}

func main() {

	fmt.Println(decodeWays("02"))
}

// Đọc lại cách 2
func decodeWays2(digits string) int {
	var dfs func(startIndex int) int
	dfs = func(startIndex int) int {
		if startIndex == len(digits) {
			return 1
		}
		ways := 0
		// can't decode string with leading 0
		if digits[startIndex] == '0' {
			return ways
		}
		// decode one digit
		ways += dfs(startIndex + 1)
		// decode two digits
		if startIndex+1 < len(digits) {
			twoDigits, _ := strconv.Atoi(digits[startIndex : startIndex+2])
			if 10 <= twoDigits && twoDigits <= 26 {
				ways += dfs(startIndex + 2)
			}
		}
		return ways
	}

	return dfs(0)
}
