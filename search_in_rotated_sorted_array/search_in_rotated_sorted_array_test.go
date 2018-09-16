package search_in_rotated_sorted_array

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := [][]int{
		{4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
	}
	testsTarget := []int{
		0,
		3,
		1,
	}
	resuts := []int{
		4,
		-1,
		5,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if res := search(tests[i], testsTarget[i]); res != resuts[i] {
			t.Fatalf("%d, actual %d, expect %d\n", i, res, resuts[i])
		}

	}

}
