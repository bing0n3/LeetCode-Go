/**
8. String to Integer (atoi)
**/
package string_to_integer

func myAtoi(str string) int {
	const (
		MaxInt32 = 2147483647
		MinInt32 = -2147483648
	)
	rst := 0
	postive := 1
	length := len(str)
	index := 0
	// runes := []rune(str)
	for index < length && str[index] == ' ' {
		index++
	}

	// if the first charter is "-" set postive to flase
	if index < length {
		if index < length && str[index] == '-' {
			postive = -1
			index++
		} else if index < length && str[index] == '+' {
			postive = 1
			index++
		}
	}

	for index < length && str[index] >= '0' && str[index] <= '9' {
		rst = rst*10 + int(str[index]-'0')
		index++
		// 这部分必须放在for循环内。因为输入的数字可能超过int, 这样数字会变为负数
		if (postive * rst) > MaxInt32 {
			return MaxInt32
		}

		if (postive * rst) < MinInt32 {
			return MinInt32
		}
	}

	return postive * rst
}
