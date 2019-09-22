package dp

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    
    cur, res := 0, 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] > nums[i-1] {
            cur++
        } else {
            cur = 1
        }
        if cur > res {
            res = cur 
        }
    }
    
    return res
}
