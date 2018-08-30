/**
* 18. 4Sum
* Time complexity O(n^3), Space O(n)
**/

package four_sum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1
			for k < l {
				if k > j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum > target {
					l--
				} else if sum < target {
					k++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
				}
			}
		}
	}
	return res
}
