// Package FinAllAnagramsInStr provides ...
package FinAllAnagramsInStr

func findAnagrams(s string, p string) []int {

	count := make([]int, 26)

	matches := 0
	var (
		LenP = len(p)
		LenS = len(s)
	)

	for i := range p {
		count[p[i]-'a'] += 1
	}

	ans := []int{}
	left, right := 0, 0

	for right < LenS {
		count[s[right]-'a']--

		if (count[s[right]-'a']) >= 0 {
			matches++
		}

		if right-left+1 > LenP {
			count[s[left]-'a']++
			if (count[s[left]-'a']) > 0 {
				matches--
			}
			left++
		}
		if matches == LenP {
			ans = append(ans, left)
		}
		right++
	}

	return ans

}
