package main

import (
	"fmt"
	"math"
	"strconv"
)

// Input: n = 3, k = 7
// Output: [181,292,707,818,929] | 3
// Explanation: Note that 070 is not a valid number, because it has leading zeroes.

// 1
// 8
// 1

func numsSameConsecDiff(n int, k int) []int {
	var result []int

	var dfs func(num int)
	dfs = func(num int) {
		if len(strconv.Itoa(num)) == n {
			result = append(result, num)
			return
		}

		last := num % 10
		for i := 0; i <= 9; i++ {
			if num != 0 && int(math.Abs(float64(last-i))) != k {
				continue
			}

			if i == 0 && num == 0 {
				continue
			}

			dfs(num*10 + i)
		}
	}

	dfs(0)

	return result
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
}
