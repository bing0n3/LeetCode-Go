/**
 * 5. Longest Palindromic Substring
 */
package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testSet := []string{
		"babad",
		"cbbd",
	}
	rstSet := []string{
		"bab",
		"bb",
	}

	caseNum := 2
	for i := 0; i < caseNum; i++ {
		if testRst := longestPalindrome(testSet[i]); testRst != rstSet[i] {
			t.Fatalf("case %d failed\nactual: %s, except: %s 。\n", i, testRst, rstSet[i])
		}
	}
}
