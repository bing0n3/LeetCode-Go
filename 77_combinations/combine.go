package combineations

func combine(n int, k int) [][]int {
	res := [][]int{}
	find(&res, &[]int{}, n, k, 1)
	return res
}

func find(res *[][]int, tmp *[]int, n, k, index int) {
	if k == 0 {
		*res = append(*res, append([]int{}, *tmp...))
	}

	for index <= n {
		*tmp = append(*tmp, index)
		find(res, tmp, n, k-1, index+1)
		*tmp = (*tmp)[:len(*tmp)-1]
		index++
	}
}
