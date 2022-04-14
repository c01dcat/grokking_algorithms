package main

import "fmt"

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	var stations = make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	stationsKey := []string{"kone", "ktwo", "kthree", "kfour", "kfive"}

	var finalStations []string

	for len(statesNeeded) > 0 {
		var bestStation string
		var statesCovered []string

		for _, station := range stationsKey {
			states := stations[station]
			var covered = intersection(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}

		}
		statesNeeded = difference(statesNeeded, statesCovered)
		finalStations = append(finalStations, bestStation)
	}
	fmt.Println(finalStations)
}

func intersection(a, b []string) []string {
	var result []string

	for _, val := range a {
		if contains(b, val) {
			result = append(result, val)
		}
	}

	return result
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

func difference(a, b []string) []string {
	var result []string

	for _, val := range a {
		if !contains(b, val) {
			result = append(result, val)
		}
	}

	return result
}