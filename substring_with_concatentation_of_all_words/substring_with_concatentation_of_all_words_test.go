/**
 * 30. Substring with Concatenation of All Words
 */
package substring_with_concatentation_of_all_words

import (
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests_s := []string{
		"barfoothefoobarman",
		"wordgoodstudentgoodword",
		"barfoofoobarthefoobarman",
		"lingmindraboofooowingdingbarrwingmonkeypoundcake",
	}
	tests_words := [][]string{
		{"foo", "bar"},
		{"word", "student"},
		{"bar", "foo", "the"},
		{"fooo", "barr", "wing", "ding", "wing"},
	}
	results := [][]int{
		{0, 9},
		{},
		{6, 9, 12},
		{13},
	}

	caseNum := len(tests_s)
	for i := 0; i < caseNum; i++ {
		res := findSubstring(tests_s[i], tests_words[i])
		if len(res) != len(results[i]) {
			t.Fatalf("CaseNum: %d, Actual: %d, Expect %d\n", i, res, results[i])
			break
		}
		for j, a := range results[i] {
			// fmt.Println(a)
			if a != res[j] {
				t.Fatalf("CaseNum: %d, Actual: %d, Expect %d\n", i, res, results[i])
				break
			}
		}
	}
}
