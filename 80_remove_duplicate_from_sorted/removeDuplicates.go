package removeduplicatefromsortedii

func removeDuplicates(nums []int) int {
	count := 0
	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[j] {
			count++
		}
		if nums[i] == nums[j] && count == 2 {
			count--
			continue
		} else if nums[i] != nums[j] {
			count = 0
		}
		j++
		nums[i], nums[j] = nums[j], nums[i]
	}

	return j + 1
}
