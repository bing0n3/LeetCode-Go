package dp

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	a := 0
	b := nums[0]

	for i := 1; i < len(nums); i++ {
		a, b = b, max(a+nums[i], b)
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
