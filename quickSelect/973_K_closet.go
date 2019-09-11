// Package KColoset provides ...
package KColoset

import "sort"

type Points [][]int

func (p Points) Less(i, j int) bool {
	if p[i][0]*p[i][0]+p[i][1]*p[i][1] < p[j][0]*p[j][0]+p[j][1]*p[j][1] {
		return true
	}

	return false
}

func (p Points) Len() int {
	return len(p)
}

func (s Points) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func kClosest(points [][]int, K int) [][]int {

	sort.Sort(Points(points))

	return points[:K]

}
