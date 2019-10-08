package bsearch

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	n := len(nums)
	lo, hi := 0, n-1
	left, right := -1, -1

	for lo < hi {
		mid := lo + (hi-lo)/2

		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if nums[lo] == target {
		left = lo
	} else {
		left = -1
	}

	hi = n - 1
	for lo < hi {
		mid := lo + (hi-lo)/2 + 1
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid
		}

	}

	if nums[lo] == target {
		right = hi
	} else {
		right = -1
	}

	return []int{left, right}

}
