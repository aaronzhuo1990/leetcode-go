package main

import (
	"fmt"
	"sort"
)

type interval struct {
	Start int
	End   int
}

func main() {
	intervals := []interval{
		{
			1,
			3,
		},
		{
			2,
			6,
		},
		{
			8,
			10,
		},
		{
			15,
			18,
		},
	}

	result := mergeIntervals(intervals)

	// Should be  [{1 6} {8 10} {15 18}]
	fmt.Printf("result %v\n", result)
}

func mergeIntervals(intervals []interval) []interval {
	result := make([]interval, 0)

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	tmp := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if tmp.End >= intervals[i].Start {
			tmp.End = max(tmp.End, intervals[i].End)
		} else {
			result = append(result, tmp)
			tmp = intervals[i]
		}
	}

	// We need to add the last interval.
	result = append(result, tmp)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
