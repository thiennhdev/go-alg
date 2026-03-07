package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int { return min(a, min(b, c)) }

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	var dp func(i int) int
	dp = func(i int) int {
		if i >= n {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		// 1-day pass
		c1 := costs[0] + dp(i+1)
		// 7-day pass
		j := i
		limit7 := days[i] + 7
		for j < n && days[j] < limit7 {
			j++
		}
		c7 := costs[1] + dp(j)
		// 30-day pass
		k := i
		limit30 := days[i] + 30
		for k < n && days[k] < limit30 {
			k++
		}
		c30 := costs[2] + dp(k)
		memo[i] = min3(c1, c7, c30)
		return memo[i]
	}
	return dp(0)
}

func main() {
	fmt.Println(mincostTickets([]int{
		1, 4, 6, 7, 8, 20,
	}, []int{
		2, 7, 15,
	}))
}
