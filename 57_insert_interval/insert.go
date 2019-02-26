package insertinterval

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return append([]Interval{}, newInterval)
	}

	flag := false
	for i := 0; i < len(intervals); i++ {
		if intervals[i].Start >= newInterval.Start {
			intervals = append(intervals, Interval{})
			copy(intervals[i+1:], intervals[i:])
			intervals[i] = newInterval
			flag = true
			break
		}
	}

	if flag == false {
		intervals = append(intervals, newInterval)
	}

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
