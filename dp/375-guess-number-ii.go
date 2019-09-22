package dp
func getMoneyAmount(n int) int {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    
    // top to down
    var helper func(int, int)int
    
    helper = func(lo, hi int) int{
        if lo >= hi {
            return 0
        }
        if dp[lo][hi] != 0 {
            return dp[lo][hi]
        }
        
        minCost := int(^uint(0) >> 1)
        for i := lo; i < hi; i++ {
            minCost = min(minCost, (i+1) + max(helper(lo,i-1), helper(i+1, hi)))
        }
        
        dp[lo][hi] = minCost
        
        return minCost
    }
    
    helper(0, n-1)
    
    return dp[0][n-1]
}


func min(a, b int) int {
    if a < b{
        return a
    }
    return b
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

