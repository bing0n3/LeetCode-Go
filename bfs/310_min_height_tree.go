package minHT

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	//build graph
	adj := make(map[int]map[int]struct{})

	for i := 0; i < len(edges); i++ {
		if adj[edges[i][0]] == nil {
			adj[edges[i][0]] = make(map[int]struct{})
		}
		if adj[edges[i][1]] == nil {
			adj[edges[i][1]] = make(map[int]struct{})
		}
		adj[edges[i][0]][edges[i][1]] = struct{}{}
		adj[edges[i][1]][edges[i][0]] = struct{}{}
	}

	leaves := []int{}

	for key, val := range adj {
		if len(val) == 1 {
			leaves = append(leaves, key)
		}
	}

	// bfs
	for n > 2 {
		n -= len(leaves)
		newLeaf := []int{}
		for _, i := range leaves {
			var j int
			for key := range adj[i] {
				j = key
			}

			_, ok := adj[j][i]
			if ok {
				delete(adj[j], i)
			}

			if len(adj[j]) == 1 {
				newLeaf = append(newLeaf, j)
			}
		}
		leaves = newLeaf
	}
	return leaves
}
