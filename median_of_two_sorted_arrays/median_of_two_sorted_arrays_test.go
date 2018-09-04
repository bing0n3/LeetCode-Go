package median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := [][][]int{
		{
			{1},
			{2, 3, 4, 5, 6},
		},
		{
			{2, 3, 4, 5, 6},
			{1},
		},
	}

	results := []float64{
		3.5,
		3.5,
	}

	caseNum := 2

	for i := 0; i < caseNum; i++ {
		if res := findMedianSortedArrays(tests[i][0], tests[i][1]); res != results[i] {
			t.Fatalf("Actual %f, Expect %f\n", res, results)
		}
	}

}
