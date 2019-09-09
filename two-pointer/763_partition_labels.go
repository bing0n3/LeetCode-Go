// Package  partitionLabels provides ...
package partitionLabels

func partitionLabels(S string) []int {

	ans := []int{}

	// create map to store char index
	last := make(map[rune]int)

	// find all letter last one
	for index, val := range S {
		last[val] = index
	}

	j, anchor := 0, 0
	for i, c := range S {
		j = max(j, last[c])
		if i == j {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
