package dp

// divede into two subproblem  

func rob(nums []int) int {
    if len(nums) ==0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    leftMax := helper(nums[1:])
    rightMax := helper(nums[:len(nums) - 1])
    
    return max(leftMax, rightMax)
}

func helper(nums []int)int {
    dp := make([]int, len(nums))
    
    dp[0] = nums[0]
    dp[1] = max(dp[0], nums[1])
    
    for i := 2; i < len(nums); i++ {
        dp[i] = max(dp[i-2]+nums[i],dp[i-1])
    }
    
    return dp[len(nums) - 1]
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
