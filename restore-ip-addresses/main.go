package main

import (
	"fmt"
	"strconv"
	"strings"
)

// A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

// For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
// Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.

// Example 1:

// Input: s = "25525511135"
// Output: ["255.255.11.135","255.255.111.35"]
// Example 2:

// Input: s = "0000"
// Output: ["0.0.0.0"]
// Example 3:

// Input: s = "101023"
// Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)

	var dfs func(startIndex int, s string, path []string)
	dfs = func(startIndex int, s string, path []string) {

		if len(path) >= 4 {
			if len(s) == startIndex { // Check all string -> valid
				result = append(result, strings.Join(path, "."))
			}
			return
		}
		// Get rest string
		for i := startIndex + 1; i <= len(s); i++ {
			if len(s[startIndex:i]) > 1 && string(s[startIndex]) == "0" {
				continue
			}

			checkNum, err := strconv.Atoi(s[startIndex:i])
			if err != nil {
				continue
			}

			if checkNum > 255 {
				continue
			}
			dfs(i, s, append(path, string(s[startIndex:i])))
		}
	}

	dfs(0, s, []string{})
	return result
}

func main() {
	fmt.Println("restoreIpAddresses: ", restoreIpAddresses("25525511135"))
}

// 1 01023
// 1 0 1023
// 1 0 1 023
