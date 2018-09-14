package longest_valid_parenthese

func longestValidParentheses(s string) int {

	stack := []int{
		-1,
	}
	maxLength := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLength = maxInt(maxLength, i-stack[len(stack)-1])
			}
		}
	}
	return maxLength
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
