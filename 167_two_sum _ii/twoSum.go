package twosumii

func twoSum(numbers []int, target int) []int {
	dict := map[int]int{}

	for i, v := range numbers {
		if t, ok := dict[v]; ok {
			return []int{t + 1, i + 1}
		} else {
			dict[target-v] = i
		}
	}
	return []int{}
}

func twoSumii(numbers []int, target int) []int {
	dict := map[int]int{}

	for i, v := range numbers {
		if t, ok := dict[v]; ok {
			return []int{t + 1, i + 1}
		} else {
			dict[target-v] = i
		}
	}
	return []int{}
}
