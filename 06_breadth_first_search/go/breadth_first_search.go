package main

import "fmt"

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	breadthFirstSearch(graph, "you")
}

func breadthFirstSearch(graph map[string][]string, name string) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[name]...)
	var searched []string

	for len(searchQueue) > 0 {
		var person = searchQueue[0]
		searchQueue = searchQueue[1:]
		personSearched := false

		for _, p := range searched {
			if p == person {
				personSearched = true
				break
			}
		}

		if !personSearched {
			if personIsSeller(person) {
				fmt.Println(person + " is a mango seller!")
				return true
			}
			searched = append(searched, person)
			searchQueue = append(searchQueue, graph[person]...)
		}
	}
	return false
}

func personIsSeller(name string) bool {
	return string(name[len(name)-1]) == "m"
}
