package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 3))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, -1))
}

func binarySearch(list []int, i int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return mid
		} else if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
