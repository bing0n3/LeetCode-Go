package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	path := []int{}
	backtracking(&res, path, nums, 0)

	return res
}

func backtracking(res *[][]int, path []int, nums []int, index int) {
	*res = append(*res, path)
	for i := index; i < len(nums); i++ {
		if i == index || nums[i] != nums[i-1] {
			path = append(path, nums[i])
			backtracking(res, append([]int{}, path...), nums, i+1)
			path = path[:len(path)-1]
		}
	}
}
