package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	write := 0
	count := 0

	for read := 0; read < n; read++ {
		count++

		if read+1 == n || chars[read+1] != chars[read] {
			chars[write] = chars[read]
			write++
			if count > 1 {
				s := strconv.Itoa(count)
				for i := 0; i < len(s); i++ {
					chars[write] = s[i]
					write++
				}
			}
			count = 0
		}
	}
	return write
}

func main() {
	fmt.Println(compress([]byte{'a', 'b', 'c'}))                                                   // 3
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))                               // 6
	fmt.Println(compress([]byte{'a'}))                                                             // 1
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'})) // 4
}
