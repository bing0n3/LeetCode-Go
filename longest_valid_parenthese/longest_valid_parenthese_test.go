package longest_valid_parenthese

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []string{
		"(()",
		"()(()",
	}
	results := []int{
		2,
		2,
	}
	caseNum := len(tests)

	for index := 0; index < caseNum; index++ {
		if res := longestValidParentheses(tests[index]); res != results[index] {
			t.Fatalf("casenum %d, actual %d, expect %d", index, res, results[index])
		}

	}
}
