package dp_comp
/**
 * @param values: a vector of integers
 * @return: a boolean which equals to true if the first player will win
 */
func firstWillWin (values []int) bool {
    // write your code here
    n := len(values)
    if n == 0 {
        return false
    }
    
    if n <= 2 {
        return true
    }
    
    dp := make([]int, n+1)
    
    dp[0] = 0
    dp[1] = values[n-1]
    dp[2] = values[n-1] + values[n-2]
    
    sum := make([]int, n + 1)
    
    for i := 1; i <= n; i++ {
        sum[i] += (sum[i-1] + values[n-i])
    }

    for i := 3; i <= n; i++ {
        dp[i] = sum[i] - min(dp[i-1],dp[i-2])
    }
    
    return 2 * dp[n] > sum[n]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
