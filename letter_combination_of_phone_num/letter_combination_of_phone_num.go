/**
 * 17. Letter Combinations of a Phone Number
**/

package letter_combination_of_phone_num

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	numLettersMap := map[rune][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	res := make([]string, 0)
	for _, num := range digits {
		if len(res) == 0 {
			for _, letter := range numLettersMap[num] {
				fmt.Println(string(letter))
				res = append(res, string(letter))
			}
			// fmt.Println(len(res))
			continue
		}
		// fmt.Println(1)
		tmp := make([]string, 0)
		// fmt.Println(len(res))
		for i := 0; i < len(res); i++ {
			for _, letter := range numLettersMap[num] {
				str := res[i]
				str += string(letter)
				tmp = append(tmp, str)
			}
		}
		res = tmp
	}

	return res
}
