package course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	// initial adjacency lis
	adjacency := make(map[int][]int)
	visited := make(map[int]int)
	incEdgeCount := []int{}

	for i := 0; i < numCourses; i++ {
		visited[i] = 0
		adjacency[i] = []int{}
		incEdgeCount = append(incEdgeCount, 0)
	}

	for _, pair := range prerequisites {
		adjacency[pair[1]] = append(adjacency[pair[1]], pair[0])
		incEdgeCount[pair[0]]++
	}

	order := []int{}
	// queue
	toVisit := []int{}

	// find the zero incoming count node
	for i := 0; i < numCourses; i++ {
		if incEdgeCount[i] == 0 {
			toVisit = append(toVisit, i)
		}
	}

	for len(toVisit) != 0 {
		from := toVisit[0]
		toVisit = toVisit[1:]
		order = append(order, from)

		for _, to := range adjacency[from] {
			incEdgeCount[to]--
			if incEdgeCount[to] == 0 {
				toVisit = append(toVisit, to)
			}
		}
	}
	if len(order) == numCourses {
		return order
	}
	return []int{}
}
