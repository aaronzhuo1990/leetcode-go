package remove_duplicates_from_sorted_array

import "sort"

func threeSum(nums []int) [][]int {

	// Here is the tricky part, sort the array before working on it!
	sort.Ints(nums)

	res := make([][]int, 0)

	for left := range nums {

		if left > 0 && nums[left] == nums[left-1] {
			// Skip the elements which have the same value.
			continue
		}

		middle, right := left+1, len(nums)-1

		for middle < right {
			s := nums[left] + nums[middle] + nums[right]
			switch {
			case s < 0:
				middle++
			case s > 0:
				right--
			default:
				res = append(res, []int{nums[left], nums[middle], nums[right]})
				// Need to move middle and right to next right position.
				middle, right = next(nums, middle, right)
			}
		}
	}

	return res
}

// next moves middle and right their next position
func next(nums []int, middle, right int) (int, int) {
	for middle < right {
		switch {
		case nums[middle] == nums[middle+1]:
			middle++
		case nums[right] == nums[right-1]:
			right--
		default:
			middle++
			right--
			return middle, right
		}
	}

	return middle, right
}
