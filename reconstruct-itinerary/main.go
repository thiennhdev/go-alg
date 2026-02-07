package main

import (
	"fmt"
)

// Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
// Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]

func InsertSorted(arr []string, v string) []string {
	res := make([]string, 0, len(arr)+1)
	inserted := false

	for _, s := range arr {
		if !inserted && v <= s {
			res = append(res, v)
			inserted = true
		}
		res = append(res, s)
	}

	if !inserted {
		res = append(res, v)
	}

	return res
}

func findItinerary(tickets [][]string) []string {
	ticketMaps := map[string][]string{}
	for _, t := range tickets {
		from, to := t[0], t[1]
		ticketMaps[from] = InsertSorted(ticketMaps[from], to)
	}

	res := make([]string, 0, len(tickets)+1)

	var dfs func(string)
	dfs = func(curr string) {
		for len(ticketMaps[curr]) > 0 {
			v := ticketMaps[curr][0]
			ticketMaps[curr] = ticketMaps[curr][1:]
			dfs(v)
		}
		res = append(res, curr)
	}

	dfs("JFK")

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main() {
	fmt.Println("findItinerary 1 : ", findItinerary(
		[][]string{
			{"JFK", "SFO"},
			{"JFK", "ATL"},
			{"SFO", "ATL"},
			{"ATL", "JFK"},
			{"ATL", "SFO"},
		},
	))

	fmt.Println("findItinerary 2 : ", findItinerary(
		[][]string{
			{"MUC", "LHR"},
			{"JFK", "MUC"},
			{"SFO", "SJC"},
			{"LHR", "SFO"},
		},
	))
}
