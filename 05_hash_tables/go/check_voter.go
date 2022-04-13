package main

import "fmt"

var voted = map[string]bool{}

func main() {
	checkVoter("tom")
	checkVoter("mike")
	checkVoter("mike")
}

func checkVoter(name string) {
	if voted[name] {
		fmt.Println("kick them out")
	} else {
		voted[name] = true
		fmt.Println("let them vote")
	}
}
