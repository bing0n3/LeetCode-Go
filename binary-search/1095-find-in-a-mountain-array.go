/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, mountainArr *MountainArray) int {
	lo, hi := 0, mountainArr.length()-1

	// find the peak

	for lo < hi {
		mid := lo + (hi-lo)/2

		if mountainArr.get(mid) > mountainArr.get(mid+1) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	peak := lo

	// find from the right
	lo = 0
	hi = peak

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mountainArr.get(mid) > target {
			hi = mid - 1
		} else if mountainArr.get(mid) < target {
			lo = mid + 1
		} else {
			return mid
		}
	}

	lo = peak
	hi = mountainArr.length() - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mountainArr.get(mid) > target {
			lo = mid + 1
		} else if mountainArr.get(mid) < target {
			hi = mid - 1
		} else {
			return mid
		}
	}

	return -1

}


