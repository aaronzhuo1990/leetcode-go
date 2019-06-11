package main

func findFirstAndLastTargetPosition(nums []int, target int) []int {

	// Find the first index of target
	index := binarySearch(nums, target)

	if index == -1 {
		return []int{-1, -1}
	}

	// Find the left index
	left := index
	for {
		i := binarySearch(nums[:left], target)

		if i == -1 {
			break
		}
		left = i
	}

	// FInd the right index
	right := index

	for {
		i := binarySearch(nums[right+1:], target)

		if i == -1 {
			break
		}

		right += 1 + i
	}

	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default:
			return mid
		}
	}

	return -1
}
