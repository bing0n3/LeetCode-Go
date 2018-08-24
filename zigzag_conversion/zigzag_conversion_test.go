/**
 * 6. ZigZag Conversion
 */
package zigzag_conversion

import "testing"

func TestConvert(t *testing.T) {
	testSet := []string{
		"PAYPALISHIRING",
		"PAYPALISHIRING",
	}
	rowNumSet := []int{
		3,
		4,
	}
	rstSet := []string{
		"PAHNAPLSIIGYIR",
		"PINALSIGYAHRPI",
	}
	caseNum := 2
	for i := 0; i < caseNum; i++ {
		if testRst := convert(testSet[i], rowNumSet[i]); testRst != rstSet[i] {
			t.Fatalf("case %d failed\nactual: %s, except: %s 。\n", i, testRst, rstSet[i])
		}
	}

}
