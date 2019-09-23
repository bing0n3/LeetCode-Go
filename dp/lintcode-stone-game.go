/**
 * @param A: An integer array
 * @return: An integer
 */
package dp
func stoneGame (A []int) int {
    // write your code here
    n := len(A)
        
        
    sum := make([][]int, n)
    dp := make([][]int, n)
    
    for i := 0; i < n; i++{
        sum[i] = make([]int, n)
        dp[i] = make([]int, n)
        sum[i][i] = A[i]
    }
    
    for l := 2; l <= n; l++ {
        for i := 0; i < n - l + 1; i++ {
            j := i + l - 1
            sum[i][j] += sum[i][j-1] + A[j]
            minVal := int(^uint(0)>>1)
            
            
            for k := i; k < j; k++ {
                if dp[i][k] + dp[k+1][j] < minVal {
                    minVal = dp[i][k] + dp[k+1][j]
                }
            }

            dp[i][j] = minVal + sum[i][j]
        }
    }
    
    return dp[0][n-1]   
}
