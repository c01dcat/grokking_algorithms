package main

import "fmt"

func main() {
	fmt.Println(recursiveCount([]int{1, 2, 3, 4}))
}

func recursiveCount(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return 1 + recursiveCount(nums[1:])
}
