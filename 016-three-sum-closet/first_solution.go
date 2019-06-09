package remove_duplicates_from_sorted_array

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	// Here is the tricky part, sort the array before working on it!
	sort.Ints(nums)

	res, delta := 0, math.MaxInt64

	for left := range nums {

		if left > 0 && nums[left] == nums[left-1] {
			// Skip the elements which have the same value.
			continue
		}

		middle, right := left+1, len(nums)-1

		for middle < right {
			s := nums[left] + nums[middle] + nums[right]
			if s < target {
				middle++
				if delta > target-s {
					delta = target - s
					res = s
				}
			} else if s > target {
				right--
				if delta > s-target {
					delta = s - target
					res = s
				}
			} else {
				return s
			}
		}
	}

	return res
}
