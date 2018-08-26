/**
 * 13. Roman to Integer
 */

package roman_to_integer

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []string{
		"III",
		"IV",
		"IX",
		"LVIII",
		"MCMXCIV",
	}
	results := []int{
		3,
		4,
		9,
		58,
		1994,
	}
	caseNum := 5

	for i := 0; i < caseNum; i++ {
		if res := romanToInt(tests[i]); res != results[i] {
			t.Fatalf("Case: %d, actual: %d, expect: %d\n", i, res, results[i])
		}
	}
}
