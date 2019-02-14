package Group_Anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	ans := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		key := string(bytes)
		if _, ok := ans[key]; ok {
			ans[key] = append(ans[key], str)
		} else {
			ans[key] = make([]string, 0)
			ans[key] = append(ans[key], str)
		}
	}
	fianlAns := make([][]string, 0, len(ans))
	for _, arr := range ans {
		fianlAns = append(fianlAns, arr)
	}

	return fianlAns
}
