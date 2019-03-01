package climbstaris

func climbStairs(n int) int {
	steps := []int{0, 1, 2}

	for i := 3; i <= n; i++ {
		steps = append(steps, steps[i-2]+steps[i-1])
	}
	return steps[n]
}
