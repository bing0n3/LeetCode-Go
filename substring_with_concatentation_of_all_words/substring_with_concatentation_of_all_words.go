/**
 * 30. Substring with Concatenation of All Words
 */
package substring_with_concatentation_of_all_words

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 || len(words)*len(words[0]) > len(s) {
		return []int{}
	}
	wordsDic := make(map[string]int)
	for _, word := range words {
		if _, ok := wordsDic[word]; ok {
			wordsDic[word]++
		} else {
			wordsDic[word] = 1
		}
	}
	substringCount := make([]int, 0)
	wordCounts := len(words)
	wordLength := len(words[0])
	conLen := wordCounts * wordLength
	for pos := 0; pos+conLen-1 < len(s); pos++ {
		seen := make(map[string]int)
		curr := pos
		for i := 0; i < wordCounts; i++ {
			// fmt.Println(pos)
			counter, ok := wordsDic[s[curr:curr+wordLength]]
			if ok {
				if _, ok := seen[s[curr:curr+wordLength]]; ok {
					seen[s[curr:curr+wordLength]]++
				} else {
					seen[s[curr:curr+wordLength]] = 1
				}
				if seen[s[curr:curr+wordLength]] > counter {
					// pos = curr - 1
					break
				} else {
					curr += wordLength
				}
			} else {
				break
			}
		}
		// fmt.Println(curr, pos)
		if (curr - pos) == conLen {
			substringCount = append(substringCount, curr-conLen)
		}
	}
	// fmt.Println(substringCount)
	return substringCount
}
