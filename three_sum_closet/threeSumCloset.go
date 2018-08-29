/**
 * 16. 3Sum Closest
 * Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
 */
package threesumcloset

import (
	"math"
	"sort"
)

func threeSum(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(target-res)) > math.Abs(float64(target-sum)) {
				res = sum
				if res == target {
					return res
				}
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
