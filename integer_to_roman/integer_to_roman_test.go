/**
 * 12. Integer to Roman
 */

package integer_to_roman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []int{
		3,
		4,
		9,
		58,
		1994,
	}
	results := []string{
		"III",
		"IV",
		"IX",
		"LVIII",
		"MCMXCIV",
	}
	caseNum := 5

	for i := 0; i < caseNum; i++ {
		if res := intToRoman(tests[i]); res != results[i] {
			t.Fatalf("Case %d,intput %d, expect: %s, actual: %s\n", i, tests[i], results[i], res)
		}
	}
}
