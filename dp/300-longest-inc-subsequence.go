// Package dp provides ...
package dp

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)

	res := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
