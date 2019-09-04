package dfs

func canFinish(numCourses int, prerequisites [][]int) bool {
	// build adjecent list first
	graph := make(map[int][]int)
	visit := make(map[int]int)

	for _, pair := range prerequisites {
		if _, ok := graph[pair[0]]; !ok {
			graph[pair[0]] = []int{}
			visit[pair[0]] = 0
		}
		graph[pair[0]] = append(graph[pair[0]], pair[1])
	}

	// dfs
	for key, _ := range graph {
		if !dfs(graph, visit, key) {
			return false
		}
	}
	return true
}

func dfs(graph map[int][]int, visit map[int]int, i int) bool {
	if visit[i] == 1 {
		return true
	}
	if visit[i] == -1 {
		return false
	}
	visit[i] = -1

	for _, j := range graph[i] {
		if !dfs(graph, visit, j) {
			return false
		}
	}
	// mark it is work, don't have cycle
	visit[i] = 1
	return true
}
