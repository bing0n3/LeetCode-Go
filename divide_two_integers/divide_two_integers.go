/**
 * 29. Divide Two Integers
 */

package divide_two_integers

func divide(dividend int, divisor int) int {
	const (
		MaxInt32 = 2147483647
		MinInt32 = -2147483648
	)
	if divisor == 0 || (dividend == MinInt32 && divisor == -1) {
		return MaxInt32
	}
	signal := 0
	if bool2int(dividend < 0)^bool2int(divisor < 0) != 1 {
		signal = 1
	} else {
		signal = -1
	}
	dividend, divisor = abs(dividend), abs(divisor)

	res := 0
	for dividend >= divisor {
		tmp, i := divisor, 1
		for dividend >= (tmp << 1) {
			i <<= 1
			tmp <<= 1
		}
		dividend -= tmp
		res += i
	}
	if signal < 0 {
		res = -res
	}
	return res
}
func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
