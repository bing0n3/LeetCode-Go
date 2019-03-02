package sortcolor

func sortColors(nums []int) {
	n := len(nums)
	i, red, blue := 0, 0, n-1

	// like two index linked list
	for i <= blue {
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		} else {
			i++
		}

	}
}
