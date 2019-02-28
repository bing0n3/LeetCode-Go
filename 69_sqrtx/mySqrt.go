package mysqrt

// method 1: brute force
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	for i := 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		}
	}

	return 0
}

// method 2: binary search
func mySqrt_binary_search(x int) int {
	if x <= 1 {
		return x
	}

	res := 0

	left, right := 1, x

	for left <= right {
		mid := (left + right) / 2

		if mid <= x/mid {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}

	return res

}
