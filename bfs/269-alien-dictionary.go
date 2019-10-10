package bfs

import "strings"

func alienOrder(words []string) string {
	graph := make(map[int][]int)
	indegree := make(map[int]int)
	totalNode := 0

	// build graph
	for _, r := range words[0] {
		if _, ok := indegree[int(r-'a')]; !ok {
			totalNode++
			indegree[int(r-'a')] = 0
		}
	}

	for i := 0; i < len(words)-1; i++ {
		for _, r := range words[i+1] {
			if _, ok := indegree[int(r-'a')]; !ok {
				totalNode++
				indegree[int(r-'a')] = 0
			}
		}
		for j := 0; j < len(words[i]); j++ {
			from := words[i][j]
			to := words[i+1][j]
			if from == to {
				continue
			}
			graph[int(from-'a')] = append(graph[int(from-'a')], int(to-'a'))
			indegree[int(to-'a')]++
			break
		}

	}

	queue := []int{}
	visited := 0
	var sb strings.Builder

	// push all node with indegree 0 to queue
	for k, v := range indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sb.WriteByte(byte(node + 'a'))
		visited++

		for _, v := range graph[node] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}

	}

	if visited != totalNode {
		return ""
	}

	return sb.String()
}
