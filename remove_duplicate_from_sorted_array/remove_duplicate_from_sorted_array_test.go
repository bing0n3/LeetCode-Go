/**
 * 26. Remove Duplicates from Sorted Array
 */
package remove_duplicate_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := [][]int{
		{1, 1, 2},
	}
	results := []int{
		2,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if res := removeDuplicates(tests[i]); res != results[i] {
			t.Fatalf("Case Number: %d, Actual: %d, Expect: %d", i, res, results[i])
		}
	}
}
