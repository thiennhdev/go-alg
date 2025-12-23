// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func dfs(nums []string, phoneNumbers map[string][]string, startIndex int, path []string, res *[]string) {
// 	if startIndex >= len(nums) {
// 		*res = append(*res, strings.Join(path, ""))
// 		return
// 	}
// 	numPhone := nums[startIndex]

// 	listChar := phoneNumbers[numPhone]

// 	startIndex++
// 	for _, s := range listChar {
// 		dfs(nums, phoneNumbers, startIndex, append(path, s), res)
// 	}
// }

// func letterCombinationsOfPhoneNumber(digits string) []string {
// 	phoneNumberMap := map[string][]string{
// 		"1": {},
// 		"2": {"a", "b", "c"},
// 		"3": {"d", "e", "f"},
// 		"4": {"g", "h", "i"},
// 		"5": {"j", "k", "l"},
// 		"6": {"m", "n", "o"},
// 		"7": {"p", "q", "r", "s"},
// 		"8": {"t", "u", "v"},
// 		"9": {"w", "x", "y", "z"},
// 		"0": {},
// 	}

// 	arr := strings.Split(digits, "")
// 	fmt.Println(arr)

// 	res := make([]string, 0)
// 	dfs(arr, phoneNumberMap, 0, []string{}, &res)
// 	return res
// }

// func main() {
// 	fmt.Println("letterCombinationsOfPhoneNumber: ", letterCombinationsOfPhoneNumber("2834"))
// }

// Algo monster
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var res []string
var phoneDigits string

var keyboard = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func dfs(startIndex int, path []rune) {
	if startIndex == len(phoneDigits) {
		res = append(res, string(path))
		return
	}
	nextNumber := rune(phoneDigits[startIndex])
	fmt.Println("nextNumber: ", string(nextNumber))
	for _, letter := range keyboard[nextNumber] {
		path = append(path, letter)
		dfs(startIndex+1, path)
		path = path[:len(path)-1]
	}
}

func letterCombinationsOfPhoneNumber(digits string) []string {
	res = []string{}
	phoneDigits = digits
	dfs(0, []rune{})
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	digits := scanner.Text()
	res := letterCombinationsOfPhoneNumber(digits)
	fmt.Println(strings.Join(res, " "))
}
