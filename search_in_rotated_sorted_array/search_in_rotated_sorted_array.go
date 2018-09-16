/**
* 33. Search in Rotated Sorted Array
* beacuse maxium time cost is O(log(n)), so it is a binary search
**/

package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	start, end := 0, len(nums)-1
	// find a mid point divde array
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}
