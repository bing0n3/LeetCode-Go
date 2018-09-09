/**
 * 10. Regular Expression Matching
 */
package regular_express_match

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return (isMatch(s, p[2:])) || (firstMatch && isMatch(s[1:], p)) //firstMatch && isMatch(s[1:], p) 改过程去除所有重复的
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
