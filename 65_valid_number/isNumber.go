package valid_number

func isNumber(s string) bool {

	n := len(s)
	isDigit := false
	isExponent := false
	Afterexponent := true
	isSign := false
	isDot := false

	for i, cur := range s {
		if '0' <= cur && cur <= '9' {
			isDigit = true
			Afterexponent = true
			continue
		}

		if cur == ' ' {
			if (i < n-1 && s[i+1] != ' ') && (isDigit || isSign || isExponent || isDot) {
				return false
			}
		} else if cur == 'e' {
			if isExponent || !isDigit {
				return false
			}
			isExponent = true
			Afterexponent = false

		} else if cur == '-' || cur == '+' {
			if i > 0 && s[i-1] != 'e' && s[i-1] != ' ' {
				return false
			}
			isSign = true
		} else if cur == '.' {
			if isDot || isExponent {
				return false
			}
			isDot = true

		} else {
			return false
		}
	}

	return Afterexponent && isDigit
}
