package permutationsii

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	generate(&res, nums, 0)
	return res
}

func generate(res *[][]int, nums []int, index int) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, nums...))
	}

	seen := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if seen[nums[i]] {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		generate(res, nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
		seen[nums[i]] = true
	}

}
