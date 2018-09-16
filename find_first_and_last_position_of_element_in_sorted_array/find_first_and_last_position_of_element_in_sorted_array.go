/**
 * 34. Find First and Last Position of Element in Sorted Array
 * Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * If the target is not found in the array, return [-1, -1].
 */

package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start, end := 0, len(nums)-1
	for nums[start] < nums[end] {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			if nums[start] == nums[mid] {
				end--
			} else {
				start++
			}
		}
	}

	if nums[start] == target && nums[end] == target {
		return []int{start, end}
	}
	return []int{-1, -1}
}
