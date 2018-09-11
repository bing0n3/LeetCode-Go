/**
 * 26. Remove Duplicates from Sorted Array
 */
package remove_duplicate_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0
	// start from the scond number, because if we start from first one, the first will not be counted
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[length] {
			continue
		}
		length++
		nums[length] = nums[i]
	}
	return length + 1
}
