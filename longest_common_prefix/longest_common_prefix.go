/**
 * 14. Longest Common Prefix
 */
package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	if num == 0 {
		return ""
	}

	if num == 1 {
		return strs[0]
	}

	baseStr := strs[0]
	baseLen := len(baseStr)
	res := ""
	for i := 0; i < baseLen; i++ {
		for j := 1; j < num; j++ {
			tmpStr := strs[j]
			if len(tmpStr) <= i || tmpStr[i] != baseStr[i] {
				return res
			}
		}
		res += string(baseStr[i])
	}
	return res
}
