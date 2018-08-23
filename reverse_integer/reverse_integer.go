/**
7. Reverse Integer
https://leetcode.com/problems/reverse-integer/description/
**/
package reverse_integer

func reverse(x int) int {
	const (
		MaxInt32 = 2147483647
		MinInt32 = -2147483648
	)
	result := 0
	tmp := x
	for ; tmp != 0; tmp = tmp / 10 {
		tmpRst := result*10 + tmp%10
		if tmpRst > MaxInt32 || tmpRst < MinInt32 {
			return 0
		}
		result = tmpRst
	}
	return result
}

func reverse2(x int) int {
	const (
		MaxInt32 = 2147483647
		MinInt32 = -2147483648
	)
	postive := true
	result := 0
	tmp := x
	if tmp < 0 {
		postive = false
		tmp *= -1
	}
	for tmp > 0 {
		result = result*10 + tmp%10
		tmp /= 10
	}
	if !postive {
		result *= -1
	}
	if result > MaxInt32 || result < MinInt32 {
		return 0
	}
	return result
}
