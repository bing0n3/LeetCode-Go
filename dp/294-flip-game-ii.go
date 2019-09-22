package dp
func canWin(s string) bool {
    
    
    dp := make(map[string]bool)
    
    var helper func(string) bool
    
    helper = func(v string) bool {
        if _, ok := dp[v]; ok {
            return dp[v]
        }
        
        chars := []rune(v)
        nowWin := false
        
        for i := 0; i < len(chars) - 1; i++ {
            if chars[i] == '+' && chars[i+1] == '+'{
                chars[i] = '-'
                chars[i+1] = '-'
                res := helper(string(chars))
                // 下一步的人输，说明这一步的人赢了
                if !res {
                    nowWin = true
                    break
                }
                chars[i] = '+'
                chars[i+1] = '+'
            }
        }
        
        dp[s] = nowWin
        
        return nowWin
    }
    
    return helper(s)
}


