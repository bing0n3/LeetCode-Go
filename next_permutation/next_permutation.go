/**
 * 31. Next Permutation
**/
package next_permutation

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	index := len(nums) - 1

	for index > 0 && nums[index-1] >= nums[index] {
		index--
	}
	// fmt.Println(index)
	if index == 0 {
		reverseSort(nums)
	} else {
		// fmt.Println(1)
		j := len(nums) - 1
		for ; j > index-1; j-- {
			if nums[index-1] < nums[j] {
				break
			}
		}
		nums[index-1], nums[j] = nums[j], nums[index-1]
		reverseSort(nums[index:])
	}

}

func reverseSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1 - i
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}
