package lengthofword

func lengthOfLastWord(s string) int {

	if s == "" {
		return 0
	}

	m := len(s) - 1
	for i := m; i >= 0; i-- {
		if s[m] == ' ' {
			m--
			continue
		}
		if s[i] == ' ' {
			return m - i
		}
	}

	return m + 1
}
