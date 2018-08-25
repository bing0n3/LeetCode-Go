/**
 * 12. Integer to Roman
 */

package integer_to_roman

import (
	"strings"
)

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	sb := new(strings.Builder)

	i := 0
	for num > 0 {
		for num >= values[i] {
			num -= values[i]
			sb.WriteString(strs[i])
		}
		i++
	}

	return sb.String()
}
