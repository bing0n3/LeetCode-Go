/**
 * 27. Remove Element
 */

package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}
