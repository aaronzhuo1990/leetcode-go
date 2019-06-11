package search_rotated_sorted_array

func searchInRotatedSortedArray(nums []int, target int) int {
	left, size := 0, len(nums)
	right := size - 1
	rotatedNum := calcRotatedNum(nums)

	for left <= right {
		mid := (left + right) / 2

		rotatedMid := (mid + rotatedNum) % size

		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case nums[rotatedMid] > target:
			right = mid - 1
		default:
			return rotatedMid
		}
	}

	return -1
}

func calcRotatedNum(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
