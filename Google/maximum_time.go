package main

import "fmt"

func maxTime(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r == '?' {
			switch i {
			case 0:
				if runes[1] >= '4' && runes[1] != '?' {
					runes[i] = '1'
				} else {
					runes[i] = '2'
				}
			case 1:
				if runes[0] == '2' {
					runes[i] = '3'
				} else {
					runes[i] = '9'
				}
			case 3:
				runes[i] = '5'
			case 4:
				runes[i] = '9'
			}
		}
	}

	return string(runes)
}

func main() {
	fmt.Println(maxTime("23:5?"))
	fmt.Println(maxTime("2?:22"))
	fmt.Println(maxTime("??:??"))
	fmt.Println(maxTime("0?:??"))
	fmt.Println(maxTime("1?:??")) // 19:59
	fmt.Println(maxTime("?4:??")) // 14:59
	fmt.Println(maxTime("?3:??")) // 23:59
	fmt.Println(maxTime("??:??")) // 23:59
}
