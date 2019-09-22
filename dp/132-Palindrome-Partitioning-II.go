package dp
func minCut(s string) int {
    if len(s) <= 1 {
        return 0
    }
    
    n := len([]rune(s))
    chars := []rune(s)
    isPal := make([][]bool, n)
    dp := make([]int, n)
    
    for i := 0; i < n; i++ {
        isPal[i] = make([]bool, n)
    }
    
    for i := 0; i < n; i++ {
        dp[i] = i
        for j := 0; j <= i; j++ {
            if (chars[i] == chars[j]) && (i - j < 2|| isPal[i-1][j+1]) {
                isPal[i][j] = true
                if j == 0 {
                    dp[i] = 0
                } else {
                    dp[i] = min(dp[i], dp[j-1]+1)
                }
            }
        }
    }
    
    
    return dp[n-1]
}


func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
