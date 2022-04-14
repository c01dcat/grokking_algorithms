package main

import (
	"fmt"
	"math"
)

var graph map[string]map[string]float64
var costs map[string]float64
var parents map[string]string
var processed []string

func main() {
	graph = make(map[string]map[string]float64)
	graph["start"] = map[string]float64{"a": 6, "b": 2}
	graph["a"] = map[string]float64{"fin": 1}
	graph["b"] = map[string]float64{"a": 3, "fin": 5}
	graph["fin"] = map[string]float64{}

	costs = make(map[string]float64)
	costs = map[string]float64{"a": 6, "b": 2, "fin": math.Inf(1)}

	parents = make(map[string]string)
	parents = map[string]string{"a": "start", "b": "start", "fin": ""}

	node := findLowestCostNode(costs)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for node := range neighbors {
			newCost := cost + neighbors[node]
			if costs[node] > newCost {
				costs[node] = newCost
				parents[node] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs)
	}

	fmt.Println("Costs from the start to each node:")
	fmt.Println(costs)
}

func findLowestCostNode(costs map[string]float64) string {
	lowestCost := math.Inf(1)
	var lowestCostNode string
	for node := range costs {
		cost := costs[node]
		if cost < lowestCost && !contains(processed, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
