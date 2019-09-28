package main

import (
	"fmt"
	"strings"
)

func compareString(A, B string) []int {
	a := strings.Split(A, ",")
	b := strings.Split(B, ",")

	n := len(b)

	store := [11]int{}
	ans := make([]int, n)

	for _, s := range a {
		if len(s) == 0 || s == "" {
			continue
		}
		hashmap := [26]int{}
		mindx := 25

		for _, r := range s {
			hashmap[int(r-'a')]++
			mindx = min(mindx, int(r-'a'))
		}
		store[hashmap[mindx]]++
	}

	for i := 1; i < 11; i++ {
		store[i] += store[i-1]
	}

	for i, s := range b {
		mindx := 25
		hashmap := [25]int{}

		for _, r := range s {
			hashmap[int(r-'a')]++
			mindx = min(mindx, int(r-'a'))
		}
		if hashmap[mindx] != 0 {
			ans[i] = store[hashmap[mindx]-1]
		}
	}

	return ans

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(compareString("abcd,aabc,bd", "aaa,aa"))
}
