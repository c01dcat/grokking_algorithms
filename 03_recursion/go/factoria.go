package main

import "fmt"

func main() {
	fmt.Println(fact(5))
}

func fact(x int) int {
	if x == 1 {
		return 1
	}

	return x * fact(x-1)
}