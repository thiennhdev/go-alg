package main

func rangeSumQueryImmutable(nums []int, left int, right int) int {
	if left < 0 || right >= len(nums) {
		return -1
	}
	target := 0
	for i := 0; i < len(nums); i++ {
		if i >= left && i <= right {
			target += nums[i]
		}
	}
	return target
}

func main() {

}
