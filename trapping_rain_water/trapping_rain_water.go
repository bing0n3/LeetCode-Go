/**
 * 42. Trapping Rain Water
 */
package trapping_rain_water

//  Brute force
func trap(height []int) int {
	res := 0
	size := len(height)
	for i := 1; i < size-1; i++ {
		max_left := 0
		max_right := 0
		for j := i; j >= 0; j-- {
			if max_left < height[j] {
				max_left = height[j]
			}
		}
		for j := i; j < size; j++ {
			if max_right < height[j] {
				max_right = height[j]
			}
		}
		if max_right > max_left {
			res += max_left - height[i]
		} else {
			res += max_right - height[i]
		}
	}
	return res
}

// dynamic programming
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	size := len(height)
	maxLeft := make([]int, size)
	maxRight := make([]int, size)

	maxLeft[0] = height[0]
	for i := 1; i < size; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}
	maxRight[size-1] = height[size-1]
	for j := size - 2; j >= 0; j-- {
		maxRight[j] = max(height[j], maxRight[j+1])
	}
	for cursor := 1; cursor < size-1; cursor++ {
		res += min(maxLeft[cursor], maxRight[cursor]) - height[cursor]
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
