// package main

// import "fmt"

// func countArrangement(n int) int {
// 	arr := make([]int, 0)
// 	for i := 1; i <= n; i++ {
// 		arr = append(arr, i)
// 	}
// 	var result int
// 	var dfs func(nums []int, path []int)
// 	dfs = func(nums []int, path []int) {

// 		if len(path) == n {
// 			check := true
// 			for i := 0; i < len(path); i++ {
// 				if path[i]%(i+1) != 0 && (i+1)%path[i] != 0 {
// 					check = false
// 				}
// 			}

// 			if check {
// 				result++
// 			}

// 			return
// 		}
// 		for i := 0; i < len(nums); i++ {
// 			dfs(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(path, nums[i]))
// 		}
// 	}

// 	dfs(arr, []int{})

// 	return result
// }

// // 3
// // 1 to n | 1 2 3
// // 1 | 2 3 => 1 2 3
// // 1 | 2 3 => 1 3 2

// // 2 | 1 3 => 2 1 3
// // 2 | 1 3 => 2 3 1

// // 3 | 1 2 => 3 1 2
// // 3 | 1 2 => 3 2 1

// func main() {
// 	fmt.Println("countArrangement: ", countArrangement(3))
// }

package main

import "fmt"

func countArrangement(n int) int {
	memo := make(map[int]int)

	var dfs func(pos int, mask int) int
	dfs = func(pos int, mask int) int {
		// Nếu đã xếp hết (pos > n) → đây là 1 hoán vị hợp lệ
		if pos > n {
			return 1
		}

		// Tạo key duy nhất cho trạng thái hiện tại
		key := (pos << 16) | mask

		fmt.Printf("pos=%d mask=%016b key=%d\n", pos, mask, key)

		// Nếu đã tính rồi → trả lại luôn
		if val, ok := memo[key]; ok {
			return val
		}

		count := 0
		// Thử từng số từ 1 → n
		for num := 1; num <= n; num++ {
			// Nếu num chưa dùng (bit num = 0)
			if (mask & (1 << num)) == 0 {
				// Nếu thỏa điều kiện đẹp
				if num%pos == 0 || pos%num == 0 {
					// Đánh dấu num đã dùng (set bit num)
					newMask := mask | (1 << num)
					count += dfs(pos+1, newMask)
				}
			}
		}

		// Lưu kết quả vào memo để lần sau dùng lại
		memo[key] = count
		return count
	}

	return dfs(1, 0)
}

func main() {
	fmt.Println(countArrangement(3))  // → 3
	fmt.Println(countArrangement(11)) // → chạy nhanh, không TLE
}
