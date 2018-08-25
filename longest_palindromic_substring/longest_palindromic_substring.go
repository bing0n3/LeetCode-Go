/**
 * 5. Longest Palindromic Substring
 */
package longest_palindromic_substring

func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandFromCenter(s, i, i)
		len2 := expandFromCenter(s, i, i+1)
		len := 0
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		// fmt.Println("a", i, len)
		if len > end-start+1 {
			// when len(s) is even, it helps, becase i is not the middle
			start = i - (len-1)/2
			end = i + len/2
		}

	}
	return s[start : end+1]
}

func expandFromCenter(s string, left int, right int) int {
	// go don't have while loop, go use for loop insted of while loop
	for left >= 0 && right < len(s) && (s[left] == s[right]) {
		left--
		right++
	}
	// left has minus 1, right has plus 1, so right - left will be larger 2 than the actual length, we should minus 1
	return right - left - 1
}
