// Package binarySearch provides ...
package binarySearch

func kthSmallest(matrix [][]int, k int) int {
	var (
		lo, hi = matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]
	)

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		j := len(matrix[0]) - 1
		for i := 0; i < len(matrix); i++ {
			// 因为行和列是递增的，如果当前列的某一行大于mid，说明后面的列的该行也大于，
			// 不需要遍历
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += (j + 1)
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
