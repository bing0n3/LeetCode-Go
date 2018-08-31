package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []string{
		"()",
	}
	result := []bool{
		true,
	}
	caseNum := 1

	for i := 0; i < caseNum; i++ {
		if res := isValid(tests[i]); res != result[i] {
			t.Fatalf("Case %d, Actual: %t, Expect: %t", i, res, result[i])
		}
	}
}
