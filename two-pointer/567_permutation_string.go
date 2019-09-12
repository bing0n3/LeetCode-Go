// Package PermutationInString provides ...
package PermutationInString

func checkInclusion(s1 string, s2 string) bool {
	count := make(map[byte]int)

	for i := range s1 {
		count[s1[i]] += 1
	}

	left, right := 0, 0

	for right < len(s2) {
		count[s2[right]] -= 1

		if right-left+1 > len(s1) {
			// remove the left one
			count[s2[left]]++
			left++
		}
		right++
		if isAllzero(count) {
			return true
		}

	}

	return false
}

func isAllzero(a map[byte]int) bool {
	for _, i := range a {
		if i != 0 {
			return false
		}
	}
	return true
}
