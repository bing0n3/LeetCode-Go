func minCostII(costs [][]int) int {
    if len(costs) == 0 || len(costs[0]) == 0 {
        return 0
    }
    
    dp := make([][]int, len(costs))
    
    
    minUpper := make([]int, len(costs[0]))
    
    for i := 0; i < len(costs); i++ {
        dp[i] = make([]int, len(costs[0]))
        
        leftPart := make([]int, len(costs[0]) + 1)
        rightPart := make([]int, len(costs[0]) + 1)
        
        leftPart[0] = int(^uint(0) >> 1)
        rightPart[len(costs[0])] = int(^uint(0) >> 1)
        
        
        for j := 0; j < len(costs[0]); j++ {
            dp[i][j] = minUpper[j] + costs[i][j] 
        }
        
        for j := 1; j <= len(costs[0]); j++ {
            leftPart[j] = min(dp[i][j-1], leftPart[j-1])
        }
        for j := len(costs[0]) - 1; j >= 0; j-- {
            rightPart[j] = min(dp[i][j], rightPart[j+1])
        }
        
        for j := 0; j < len(costs[0]); j++ {
            minUpper[j] = min(leftPart[j], rightPart[j+1]) 
        }
    }
    
    res := int(^uint(0) >> 1)
    for i := 0; i < len(costs[0]); i++ {
        if dp[len(costs) -1][i] < res {
            res = dp[len(costs) -1][i]
        }
    }
    
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
