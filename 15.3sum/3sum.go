/**
 * 15. 3Sum
 */

package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)
	solution := make([][]int, 0)
	sort.Ints(nums)
	for i := range nums {
		firstNum := nums[i]
		if i > 0 && firstNum == nums[i-1] {
			continue
		}
		j := i + 1
		k := length - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			secondNum := nums[j]
			thirdNum := nums[k]
			sum := firstNum + secondNum + thirdNum
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				solution = append(solution, []int{firstNum, secondNum, thirdNum})
				j++
			}
		}
	}
	return solution
}
