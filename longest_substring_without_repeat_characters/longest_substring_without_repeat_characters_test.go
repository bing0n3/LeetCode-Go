package longest_substring_without_repeat_characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testSet := []string{
		"abcabcbb",
		"aaaaaaaa",
		"abcde",
		"",
		"pwwkew",
	}
	resultSet := []int{
		3,
		1,
		5,
		0,
		3,
	}
	caseNum := 5
	for i := 0; i < caseNum; i++ {
		if length := lengthOfLongestSubstring(testSet[i]); length != resultSet[i] {
			t.Fatalf("case %d failed\nactual: %s, except: %d\n", i, testSet[i], resultSet)
		}

	}
}
