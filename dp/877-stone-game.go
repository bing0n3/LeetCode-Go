package dp

func stoneGame(piles []int) bool {
    // up to down 
    n := len(piles)
    if n <= 1 {
        return true
    }
    
    // 
    dp := make([]int, n)
    copy(dp, piles)
    
    for l := 2; l <= n; l++ {
        for i := 0; i < n - l + 1; i++ {
            dp[i] = max(piles[i]-dp[i+1], piles[i + l - 1]-dp[i])
        }
    }
    
    return dp[0] > 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}

