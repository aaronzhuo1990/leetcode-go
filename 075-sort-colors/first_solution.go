package main

import "fmt"

func main() {

	colors := []int{2, 0, 2, 1, 1, 0}

	sortColors(colors)

	fmt.Printf("colors %v\n", colors) // should be {0, 0, 1, 1, 2, 2}
}

func sortColors(colors []int) {
	left, middle, right := 0, 0, len(colors)-1

	for middle <= right {
		switch colors[middle] {
		case 0:
			colors[left], colors[middle] = colors[middle], colors[left]
			left++
			middle++
		case 1:
			middle++
		case 2:
			colors[middle], colors[right] = colors[right], colors[middle]
			right--
		}
	}
}
