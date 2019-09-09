// Package subarrayLessK provides ...
package subarrayLessK

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, prod, ans := 0, 1, 0

	for right, val := range nums {
		prod *= val
		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
