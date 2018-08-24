/**
8. String to Integer (atoi)
**/
package string_to_integer

import "testing"

func TestMyAtoi(t *testing.T) {
	testSet := []string{
		"-91283472332",
		"9223372036854775808",
		"42",
	}
	rstSet := []int{
		-2147483648,
		2147483647,
		42,
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if cRst := myAtoi(testSet[i]); cRst != rstSet[i] {
			t.Fatalf("case %d failed\nactual: %d, except: %d\n", i, cRst, rstSet[i])
		}
	}
}
