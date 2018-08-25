/**
 * 11. Container with most water
 */

package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	tests := [][]int{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
	}
	results := []int{
		49,
	}
	caseNum := 1

	for i := 0; i < caseNum; i++ {
		if ret := maxArea2(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
