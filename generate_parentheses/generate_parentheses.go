/**
 * 22. Generate Parentheses
 */

package generate_parentheses

func generateParenthesis(n int) []string {
	result := new([]string)
	length := 2 * n
	generateAll(make([]rune, n*2), 0, length, result)
	return *result
}

// recursive to generate all possible set
func generateAll(current []rune, pos int, length int, result *[]string) {
	if pos == length {
		if isValid(current) {
			*result = append(*result, string(current))
		}
	} else {
		current[pos] = '('
		generateAll(current, pos+1, length, result)
		current[pos] = ')'
		generateAll(current, pos+1, length, result)
	}
}

func isValid(current []rune) bool {
	blance := 0
	for _, char := range current {
		if char == '(' {
			blance++
		} else {
			blance--
		}
		if blance < 0 {
			return false
		}
	}
	return blance == 0
}

// track the number of brackets
func generateParenthesis2s(n int) []string {
	answer := new([]string)
	trackBackets("", 0, 0, n, answer)
	return *answer
}

func trackBackets(current string, open int, close int, n int, answer *[]string) {
	if len(current) == n*2 {
		*answer = append(*answer, current)
		return
	}

	if open < n {
		trackBackets(current+"(", open+1, close, n, answer)
	}

	if close < open {
		trackBackets(current+")", open, close+1, n, answer)
	}
}

// closure number
func generateParenthesis3(n int) []string {
	if n == 0 {
		return []string{""}
	}
	answer := make([]string, 0)
	for c := 0; c < n; c++ {
		for _, left := range generateParenthesis(c) {
			for _, right := range generateParenthesis(n - c - 1) {
				answer = append(answer, "("+left+")"+right)
			}
		}
	}
	return answer
}
