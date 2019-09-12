// Package combSumIII provides ...
package combSumIII

func combinationSum3(k int, n int) [][]int {

	res := [][]int{}

	var dfs func(int, int, []int)

	dfs = func(a int, target int, ans []int) {
		if len(ans) == k {
			if target == 0 {
				res = append(res, ans)
			}
			return
		}

		for ; a <= 9; a++ {
			if target-a >= 0 {
				tmp := make([]int, len(ans))
				copy(tmp, ans)
				tmp = append(tmp, a)
				dfs(a+1, target-a, tmp)
			}
		}
	}

	dfs(1, n, []int{})

	return res
}
