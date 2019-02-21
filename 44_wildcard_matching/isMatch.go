package wildcardmatch

// dp
func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)

	dp := make([][]bool, lenS+1)
	for i := range dp {
		dp[i] = make([]bool, lenP+1)
	}

	for i := 0; i <= lenS; i++ {
		for j := 0; j <= lenP; j++ {
			// if both empty. pattern
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}

			// if only have string, don't have pattern,
			// impossible to pattern
			if j == 0 {
				dp[i][j] = false
				continue
			}

			// if only pattern, don't have string
			if i == 0 {
				dp[0][j] = p[j-1] == '*' && dp[0][j-1]
				continue
			}

			//if *, * pattern empty substring or * have patterned some string
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && dp[i-1][j-1]
			}

		}
	}

	return dp[lenS][lenP]
}
