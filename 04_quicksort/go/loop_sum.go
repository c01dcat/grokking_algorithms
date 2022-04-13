package main

import "fmt"

func main() {
	fmt.Println(loopSum([]int{1, 2, 3, 4}))
}

func loopSum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
