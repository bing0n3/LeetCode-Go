package subsets

func subsets(nums []int) [][]int {
	ans := [][]int{[]int{}}

	for _, n := range nums {
		newSets := make([][]int, 0)
		for _, set := range ans {
			tmp := make([]int, len(set)+1)
			// append slice will effect the orignal array
			copy(tmp, append(set, n))
			newSets = append(newSets, tmp)
		}
		ans = append(ans, newSets...)
	}
	return ans
}
