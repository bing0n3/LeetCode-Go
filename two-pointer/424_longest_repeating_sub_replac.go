package longestRepeatingRplc

func characterReplacement(s string, k int) int {
	var (
		left  = 0
		count = make(map[byte]int)
	)

	max_count, result := 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]] += 1
		max_count = max(max_count, count[s[right]])
		if right-left+1-max_count > k {
			count[s[left]] -= 1
			left++
		}
		result = max(result, right-left+1)
	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
