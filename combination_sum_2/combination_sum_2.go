/**
 * 40. Combination Sum II
 * https://leetcode.com/problems/combination-sum-ii/description/
 */

package combination_sum_2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	combination(candidates, target, &res, make([]int, 0), 0)
	return res
}

func combination(candidates []int, target int, res *[][]int, exitst []int, index int) {
	if target == 0 {
		*res = append(*res, exitst)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i != 0 && candidates[i-1] == candidates[i] && i > index {
			continue
		}
		if candidates[i] <= target {
			tmp := make([]int, len(exitst))
			copy(tmp, exitst)
			tmp = append(tmp, candidates[i])
			combination(candidates, target-candidates[i], res, tmp, i+1)
		}
	}
}
