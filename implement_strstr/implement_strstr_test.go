/**
 * 28. Implement strStr()
 */
package implement_strstr

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := [][]string{
		{"hello", "ll"},
		{"aaaaa", "bba"},
	}
	results := []int{
		2,
		-1,
	}
	caseNum := len(tests)

	for i := 0; i < caseNum; i++ {
		if res := strStr(tests[i][0], tests[i][1]); res != results[i] {
			t.Fatalf("Case Num: %d, actual %d, expect %d\n", i, res, results[i])
		}
	}

}
