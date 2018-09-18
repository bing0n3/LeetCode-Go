/**
 * 46. Permutations
 * Given a collection of distinct integers, return all possible permutations.
 * https://leetcode.com/problems/permutations/description/
 */

package permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	generate_permute(nums, &result, 0)
	return result
}

func generate_permute(nums []int, res *[][]int, index int) {

	if index == len(nums) {
		tmp := make([]int, len(nums))
		// why not slice, beacuse slice creat a new slice point to original array, so
		// operate in array also operate in slice
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := index; i < len(nums); i++ {
		// backtrack
		nums[i], nums[index] = nums[index], nums[i]
		tmp := nums[:]
		generate_permute(tmp, res, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
