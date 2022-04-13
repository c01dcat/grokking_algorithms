package main

import "fmt"

func main() {
	fmt.Println(quicksort([]int{3, 2, 1, 4, 5, 6, 7, 8, 9, 10}))
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var less, greater []int
	for _, num := range nums[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	less = append(quicksort(less), pivot)
	greater = quicksort(greater)
	return append(less, greater...)
}
