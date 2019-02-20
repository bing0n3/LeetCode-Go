package largestRectangleArea

/*
two method
1. burte force. each histogram, scan to right and left.
2. Monotonic stack

this version is method 1, and do some optimze, save result to reuse.
*/

func largestRectangleArea(heights []int) int {
	length := len(heights)

	if length == 0 {
		return 0
	}
	if length == 1 {
		return heights[0]
	}
	leftArr := make([]int, length)
	rightArr := make([]int, length)

	// scan from left to int
	rightArr[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		if heights[i] > heights[i+1] {
			rightArr[i] = 1
		} else {
			j := i + 1
			for j < length && heights[j] >= heights[i] {
				j += rightArr[j]
			}
			rightArr[i] = j - i
		}
	}

	// scan from right to left
	leftArr[0] = 1
	for i := 1; i < length; i++ {
		if heights[i] > heights[i-1] {
			leftArr[i] = 1
		} else {
			j := i - 1
			for j >= 0 && heights[j] >= heights[i] {
				j -= leftArr[j]
			}
			leftArr[i] = i - j
		}
	}
	maxArea := heights[0]
	var tmpA int
	for i := 0; i < length; i++ {
		tmpA = heights[i] * (leftArr[i] + rightArr[i] - 1)
		if tmpA > maxArea {
			maxArea = tmpA
		}

	}

	return maxArea
}
