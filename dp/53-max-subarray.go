func maxSubArray(nums []int) int {
    cur := 0
    res := nums[0]
    
    for _, num := range nums {
        cur = max(num, cur + num)
        res = max(res, cur)
    }
    
    return res
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

