/**
 * 27. Remove Element
 */

package remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := [][]int{
		{3, 2, 2, 3, 3},
	}
	results := []int{
		2,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if res := removeElement(tests[i][:len(tests[i])-1], tests[i][len(tests[i])-1]); res != results[i] {
			t.Fatalf("Casu Num: %d, Actual: %d, Expect: %d", i, res, results[i])
		}
	}

}
