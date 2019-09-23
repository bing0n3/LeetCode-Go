/**
 * @param nums: A list of integers
 * @return: An integer indicate the value of maximum difference between two substrings
 */

func maxDiffSubArrays (nums []int) int {
    // write your code here
    n := len(nums)
    max_l := nums[0]
    max_l2r := make([]int, n)
    max_r := nums[n-1]
    max_r2l := make([]int, n)
    min_l := nums[0]
    min_l2r := make([]int, n)
    min_r := nums[n-1]
    min_r2l := make([]int, n)
    
    max_l2r[0] = nums[0]
    min_l2r[0] = nums[0]
    max_r2l[n-1] = nums[n-1]
    min_r2l[n-1] = nums[n-1]


    for i := 1; i < n; i++ {
        max_l = max(nums[i], max_l+nums[i])
        max_l2r[i] = max(max_l, max_l2r[i-1])
        
        min_l = min(nums[i], min_l + nums[i])
        min_l2r[i] = min(min_l, min_l2r[i-1])
        
        max_r = max(nums[n - i - 1], max_r + nums[n - i - 1])
        max_r2l[n-i-1] = max(max_r, max_r2l[n-i])
        
        min_r = min(nums[n - i - 1], min_r + nums[n - i -1])
        min_r2l[n-i-1] = min(min_r, min_r2l[n-i])
    }
    
    
    res := 0
    
    for i := 0; i < n-1; i++ {
        res = max(res,max(abs(max_l2r[i] - min_r2l[i+1]),abs(max_r2l[i+1]- min_l2r[i])))  
    }
    

    
    return res
}

func abs(a int) int {
    if a >= 0 {
        return a 
    }
    
    return -a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
