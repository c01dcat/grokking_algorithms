package main

import "fmt"

func main() {
	s := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(s))
}

func selectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
