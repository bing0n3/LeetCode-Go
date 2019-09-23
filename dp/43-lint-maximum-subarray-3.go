/**
 * @param nums: A list of integers
 * @param k: An integer denote to find k non-overlapping subarrays
 * @return: An integer denote the sum of max k non-overlapping subarrays
 */
func maxSubArray (nums []int, k int) int {
    if len(nums) == 0 || k == 0 {
        return 0
    }
    // write your code here
    localMax := make([][]int, len(nums)+1)
    globalMax := make([][]int, len(nums)+1)
    
    for i := range localMax {
        localMax[i] = make([]int, k+1)
        globalMax[i] = make([]int, k+1)
    }
    
    for i := 1; i < len(nums) + 1; i++ {
        for j := 1; j < k + 1 && j <= i; j++ {
            localMax[j-1][j] = -int(^uint(0)>>1)
            
            localMax[i][j] = max(localMax[i-1][j],globalMax[i-1][j-1]) + nums[i-1]
            if i == j {
                globalMax[i][j] = localMax[i][j]
            } else {
                globalMax[i][j] = max(localMax[i][j], globalMax[i-1][j])
            }
        }
    }
    
    return globalMax[len(nums)][k]
}

func max(a, b int) int{
    if a > b {
        return a
    }
    
    return b
}
