package next_permutation

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := [][]int{
		{1, 1},
		{1},
		// {1, 2},
		{1, 2, 3},
		{4, 2, 3, 1},
		{1, 3, 2},
		{2, 3, 1},
	}
	results := [][]int{
		{1, 1},
		{1},
		// {2, 1},
		{1, 3, 2},
		{4, 3, 1, 2},
		{2, 1, 3},
		{3, 1, 2},
	}
	caseNum := len(tests)

	for i := 0; i < caseNum; i++ {
		nextPermutation(tests[i])
		for j, num := range tests[i] {
			if num != results[i][j] {
				t.Fatalf("case num %d, actual %d, expect %d", i, tests[i], results[i])
				break
			}
		}
	}
}
