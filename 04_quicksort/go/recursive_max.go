package main

import "fmt"

func main() {
	fmt.Println(recursiveMax([]int{1, 2, 3, 4}))
}

func recursiveMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := recursiveMax(nums[1:])
	if nums[0] > max {
		return nums[0]
	}
	return max
}
