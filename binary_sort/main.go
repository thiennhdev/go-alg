package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 2 8 89 120 1000 -> 8
// 89 120 1000 -> 89
// 120 1000 ->
func binarySearch(arr []int, target int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	if len(arr) == 1 {
		return 0
	}
	n := len(arr)
	start := 0
	end := len(arr) - 1
	midIndex := n / 2

	for end > start {
		if arr[midIndex] == target {
			return midIndex
		}
		if target < arr[midIndex] {
			end = midIndex
		}

		if target > arr[midIndex] {
			start = midIndex + 1
		}

		midIndex = (end + start) / 2
	}

	return -1
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	res := binarySearch(arr, target)
	fmt.Println(res)
}
