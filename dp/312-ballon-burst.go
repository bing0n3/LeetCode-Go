package dp

func maxCoins(nums []int) int {

    n := len(nums)
    dp := make([][]int, n + 2)
    nums = append(append([]int{1}, nums...),1)    
    for i := range nums {
        dp[i] = make([]int, n + 2)
    }
    
    // add padding

    for l := 1; l <= n; l++ {
        for i := 1; i + l - 1 < n + 1; i++ {
            j :=  i + l - 1
            
            maxVal := 0
            for k := i; k <= j; k++ {
                cur := nums[i-1] * nums[k] * nums[j+1]
                maxVal = max(maxVal, cur + dp[i][k-1] + dp[k+1][j])
            }
            
            dp[i][j] = maxVal
        }
    }
    return dp[1][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}
