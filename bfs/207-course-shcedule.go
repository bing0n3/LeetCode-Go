package bfs

func canFinish(numCourses int, prerequisites [][]int) bool {
	// build DAG first
	graph := make([][]int, numCourses)
	indegree := make(map[int]int)

	// calc indegree each node and build adj list
	for _, pairs := range prerequisites {
		indegree[pairs[1]]++
		if graph[pairs[0]] != nil {
			graph[pairs[0]] = append(graph[pairs[0]], pairs[1])
		} else {
			graph[pairs[0]] = append([]int{}, pairs[1])
		}
	}

	queue := []int{}
	visitedCount := 0

	// put all node indegree into queue
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visitedCount++
		for _, v := range graph[node] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return visitedCount == numCourses

}
