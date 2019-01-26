package plusOne

func plusOne(digits []int) []int {
	m := len(digits) - 1
	tmp := 0
	for ; m >= 0; m-- {
		if digits[m] == 9 {
			digits[m] = 0
			tmp = 1
		} else {
			digits[m]++
			tmp = 0
			return digits
		}
	}
	if tmp == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
