package subsets

import "testing"

func TestSubsets(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3},
	}

	results := [][][]int{
		[][]int{[]int{}, []int{3}, []int{2}, []int{1}, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 2, 3}},
	}

	for i := 0; i < len(tests); i++ {
		if observed := subsets(tests[i]); len(observed) != len(results[i]) || true {
			t.Fatalf("observer: %v,expect: %v", observed, results[i])
		}

	}
}
