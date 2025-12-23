package main

import "fmt"

func squareRoot(n int) int {
	startPtr, endPtr := 0, n
	bound := -1
	for startPtr <= endPtr {
		midPtr := (startPtr + endPtr) / 2
		if midPtr*midPtr >= n {
			if midPtr*midPtr == n {
				return midPtr
			}
			bound = midPtr
			endPtr = midPtr - 1
		} else {
			startPtr = midPtr + 1
		}
	}
	return bound - 1
}

func main() {
	fmt.Println(squareRoot(16))
}
