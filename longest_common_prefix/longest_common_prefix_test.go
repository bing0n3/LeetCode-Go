/**
 * 14. Longest Common Prefix
 */
package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := [][]string{
		{"aa", "a"},
	}
	results := []string{
		"a",
	}
	caseNum := 1

	for index := 0; index < caseNum; index++ {
		if testRst := longestCommonPrefix(tests[index]); testRst != results[index] {
			t.Fatalf("case %d failed\nactual: %s, except: %s ã€‚\n", index, testRst, results[index])
		}
	}
}
