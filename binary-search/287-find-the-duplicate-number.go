// Package binarySearch provides ...
package binarySearch

func findDuplicate(nums []int) int {
	// binary search range
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		cnt := 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
