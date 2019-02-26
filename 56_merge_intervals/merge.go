package mergeintervals

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	var res []Interval
	cur := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if cur.End >= intervals[i].Start {
			if intervals[i].End > cur.End {
				cur.End = intervals[i].End
			}
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}

	res = append(res, cur)
	return res
}
