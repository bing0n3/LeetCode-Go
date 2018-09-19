/**
 * 39. Combination Sum
 * https://leetcode.com/problems/combination-sum/description/
 */

package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	combination(candidates, target, &res, make([]int, 0), 0)
	return res
}

func combination(candidates []int, target int, res *[][]int, exitst []int, index int) {
	if target == 0 {
		*res = append(*res, exitst)
		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] <= target {
			tmp := make([]int, len(exitst))
			copy(tmp, exitst)
			tmp = append(tmp, candidates[i])
			combination(candidates, target-candidates[i], res, tmp, i)
		}
	}
}
