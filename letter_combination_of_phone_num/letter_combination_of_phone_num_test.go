/**
 * 17. Letter Combinations of a Phone Number
**/

package letter_combination_of_phone_num

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []string{
		"23",
	}
	results := [][]string{
		{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	}
	caseNum := 1

	for i := 0; i < caseNum; i++ {
		if res := letterCombinations(tests[i]); !reflect.DeepEqual(results[i], res) {
			t.Fatalf("case %d, actul %s, expect %s", caseNum, res, results[i])
		}
	}
}
