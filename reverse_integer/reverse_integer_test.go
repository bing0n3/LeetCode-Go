package reverse_integer

import "testing"

func TestReverse2(t *testing.T) {
	testSet := []int{
		123,
		-123,
		120,
	}
	rstSet := []int{
		321,
		-321,
		21,
	}
	caseNum := 1

	for i := 0; i < caseNum; i++ {
		if cRst := reverse(testSet[i]); cRst != rstSet[i] {
			t.Fatalf("case %d failed\nactual: %d, except: %d\n", i, cRst, rstSet[i])
		}
	}
}
