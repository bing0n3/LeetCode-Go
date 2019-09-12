// Package sildingWindowMax provides ...
package sildingWindowMax

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 || (len(nums) < k) {
		return []int{}
	}

	// initial window
	var (
		ans   = make([]int, len(nums)-k+1)
		stack = []int{}
	)

	// initial first
	for i := 0; i < len(nums); i++ {
		// if out range remove element from stack
		for len(stack) != 0 && stack[0] < i-k+1 {
			stack = stack[1:]
		}

		// remove node smaller than current node
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
		if i >= k-1 {
			ans[i-k+1] = nums[stack[0]]
		}
	}

	return ans

}
