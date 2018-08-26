/**
 * 13. Roman to Integer
 */

package roman_to_integer

func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	cur, prv, out := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur = values[s[i]]
		if cur < prv {
			out -= cur
		} else {
			out += cur
		}
		prv = cur
	}
	return out
}
