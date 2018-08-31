/**
 * 20. Valid Parentheses
 */

package valid_parentheses

func isValid(s string) bool {
	stack := make([]rune, 0)
	parentheses := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, char)
		case '{':
			stack = append(stack, char)
		case '[':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || parentheses[')'] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || parentheses['}'] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || parentheses[']'] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return !(len(stack) != 0)
}
