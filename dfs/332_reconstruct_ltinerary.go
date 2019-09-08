// Package ReconstructLtinerary provides ...
package ReconstructLtinerary

import "sort"

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return []string{}
	}

	N := len(tickets)

	// construct graph
	graph := make(map[string][]string)
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}
	// sort by lexical order
	for key, _ := range graph {
		sort.Strings(graph[key])
	}

	route := make([]string, 0, len(tickets)+1)
	route = append(route, "JFK")

	var dfs func(string)
	dfs = func(from string) {
		if _, ok := graph[from]; !ok {
			// if it cannot reach oterh airport return
			return
		}

		for i := 0; i < len(graph[from]); i++ {
			tmp := make([]string, len(graph[from]))
			copy(tmp, graph[from])
			to := graph[from][i]
			route = append(route, to)
			graph[from] = append(graph[from][:i], graph[from][i+1:]...)
			dfs(to)
			if len(route) == N+1 {
				return
			}
			// revert change
			graph[from] = tmp
			route = route[:len(route)-1]
		}

	}

	dfs("JFK")

	return route
}
