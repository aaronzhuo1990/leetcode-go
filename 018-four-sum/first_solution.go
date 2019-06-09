package four_sum

import (
	"sort"
)

// fourSum finds all unique quadruplets in the array which gives the sum of target.
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	// Here is the tricky part, sort the array before working on it!
	sort.Ints(nums)

	for left := 0; left < len(nums)-3; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			// Move left to the right position to avoid duplicating the same result.
			continue
		}

		for mid1 := left + 1; mid1 < len(nums)-2; mid1++ {
			if mid1 > left+1 && nums[mid1] == nums[mid1-1] {
				// Move mid1 to the right position to avoid duplicating the same result.
				continue
			}
			mid2, right := left+1, len(nums)-1

			for mid2 < right {
				sum := nums[left] + nums[mid1] + nums[mid2] + nums[right]
				if sum < target {
					mid2++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{nums[left], nums[mid1], nums[mid2], nums[right]})
					mid2, right = next(nums, mid2, right)
				}
			}
		}

	}

	return result
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
