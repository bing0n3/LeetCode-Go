// Package water provides ...
package main

import "fmt"

func water(a []int, capacity int) int {
	if len(a) == 0 {
		return 0
	}

	res := capacity
	remain := capacity

	step := 0

	for i, need := range a {
		if remain >= need {
			remain -= need
		} else {
			res += capacity - remain
			remain = capacity - need
			step += (2 * i)
		}
		step++
	}
	return step
}

func main() {
	testCase := [][]int{
		{2, 4, 5, 1, 2},
	}
	capacity := []int{6}

	fmt.Println(water(testCase[0], capacity[0]))

}
