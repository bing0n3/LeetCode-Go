// Package MinWindowSub provides ...
package MinWindowSub

func minWindow(s string, t string) string {
	var (
		counter = make([]int, 128)
		matches = len(t)
		l, r    = 0, 0
		d       = int(^uint(0) >> 1)
		head    = -1
	)

	for i := 0; i < len(t); i++ {
		counter[t[i]]++
	}

	for r < len(s) {
		counter[s[r]]--

		if counter[s[r]] >= 0 {
			matches--
		}

		for matches == 0 {
			if r-l+1 < d {
				d = r - l + 1
				head = l
			}

			// remove left one
			counter[s[l]]++
			if counter[s[l]] > 0 {
				matches++
			}
			l++
		}

		r++
	}

	if head < 0 {
		return ""
	} else {
		return s[head : head+d]
	}
}
