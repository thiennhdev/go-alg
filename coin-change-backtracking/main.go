package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Tim xem co bao nhieu dong
// amount = 1 thi co it nhat bao nhieu dong de ra 1
// amount = 5 thi co it nhat bao nhieu dong sum lai bang 5
// memo[amount]way
func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	var minCoins func(sum int) int
	minCoins = func(sum int) int { // Return int is number of ways can sum by total
		ans := math.MaxInt

		if sum == amount {
			return 0
		}

		if sum > amount {
			return math.MaxInt
		}

		if _, ok := memo[sum]; ok {
			return memo[sum]
		}
		for _, coin := range coins {
			result := minCoins(sum + coin)
			if result == math.MaxInt {
				continue
			}
			ans = min(ans, result+1)
		}
		memo[sum] = ans
		return ans
	}
	a := minCoins(0)
	if a == math.MaxInt {
		a = -1
	}

	return a
}

// memo amount = số đồng ít nhất để tạo ra amount
func main() {
	fmt.Println(coinChange([]int{1, 5, 2}, 11))
}
