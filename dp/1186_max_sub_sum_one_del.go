// Package maxSubSumOneDel provides ...
package maxSubSumOneDel

func maximumSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	N := len(arr)
	dp := make([]int, N)
	dp_r := make([]int, N)

	dp[0] = arr[0]
	dp_r[N-1] = arr[N-1]
	res := arr[0]

	for i := 1; i < N; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + arr[i]
		} else {
			dp[i] = arr[i]
		}
		res = max(res, dp[i])
	}

	for i := N - 2; i >= 0; i-- {
		if dp_r[i+1] > 0 {
			dp_r[i] = dp_r[i+1] + arr[i]
		} else {
			dp_r[i] = arr[i]
		}
	}
	for i := 1; i < N-1; i++ {
		if arr[i] < 0 {
			res = max(res, dp[i-1]+dp_r[i+1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
