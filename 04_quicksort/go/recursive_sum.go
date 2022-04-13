package main

import "fmt"

func main() {
	fmt.Println(recursiveSum([]int{1, 2, 3, 4}))
}

func recursiveSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] + recursiveSum(nums[1:])
}
