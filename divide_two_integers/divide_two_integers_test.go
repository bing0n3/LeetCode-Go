/**
 * 29. Divide Two Integers
 */

package divide_two_integers

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := [][]int{
		{0, 1},
		{-2147483648, -1},
	}
	results := []int{
		0,
		2147483647,
	}
	caseNum := len(tests)

	for i := 0; i < caseNum; i++ {
		if res := divide(tests[i][0], tests[i][1]); res != results[i] {
			t.Fatalf("Case Num: %d, actual: %d, expect %d", i, res, results[i])
		}
	}
}
