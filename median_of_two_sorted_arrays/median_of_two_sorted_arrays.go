/**
 * 4. Median of Two Sorted Arrays
 **/

package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length&1 == 1 {
		return float64(findKth(nums1, 0, nums2, 0, length/2+1))
	} else {
		res1 := findKth(nums1, 0, nums2, 0, length/2)
		// fmt.Println("Cal 2")
		res2 := findKth(nums1, 0, nums2, 0, length/2+1)
		return float64(res1+res2) / 2
	}
}

func findKth(left []int, startL int, right []int, startR int, k int) int {

	// fmt.Println(startL)
	// fmt.Println(startR)
	// fmt.Println("K", k)
	if startL >= len(left) {
		return right[startR+k-1]
	}
	if startR >= len(right) {
		return left[startL+k-1]
	}
	if k == 1 {
		return min(left[startL+k-1], right[startR+k-1])
	}

	m1 := min(len(left)-1, startL+k/2-1)  // 1
	m2 := min(len(right)-1, startR+k/2-1) // 1

	// fmt.Println("?", startL+(k/2)-1)
	if left[m1] >= right[m2] {
		// fmt.Println("StartR", startR)
		return findKth(left, startL, right, m2+1, k-(m2-startR+1))
	} else {
		// fmt.Println("StartR", startR)
		return findKth(left, m1+1, right, startR, k-(m1-startL+1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
