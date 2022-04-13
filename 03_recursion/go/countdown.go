package main

import "fmt"

func main() {
	countdown(5)
}

func countdown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	countdown(n - 1)
}
