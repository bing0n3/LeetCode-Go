package findpeakelement


// we can solve this problem using brute force method. need time complexity O(n)
// however the probelm requirement is logarithmic
// so we can easily decide use binary search for it 
// when nums[mid] less than nums[mid+1], it means that right part must have a peak 
// when nums[mid] larger than nums[mid+1], it means that mid is possible be peak or left part must have one 
func findPeakElement(nums []int) int {
	lo, hi = 0, len(nums) - 1

	for lo < hi {
		mid = lo + (hi - lo) / 2
		if nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

