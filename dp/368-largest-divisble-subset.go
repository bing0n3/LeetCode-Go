package dp

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sort.Ints(nums)
	dp := make([]int, len(nums))
	pre := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		pre[i] = -1
	}

	max := 0
	idx := -1

	for i := 0; i < len(nums); i++ {
		dp[i], pre[i] = 1, -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				pre[i] = j
			}
		}
		if dp[i] > max {
			max, idx = dp[i], i
		}
	}

	ans := []int{}
	for ; idx != -1; idx = pre[idx] {
		ans = append(ans, nums[idx])

	}

	return ans
}
