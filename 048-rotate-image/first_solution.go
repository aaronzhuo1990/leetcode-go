package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotateImage(matrix)

	fmt.Printf("rotated matrix:\n")
	for _, line := range matrix {
		fmt.Printf("\t%v\n", line)
	}

	// Expected result:
	// [
	//	[15,13, 2, 5],
	//  [14, 3, 4, 1],
	//  [12, 6, 8, 9],
	//  [16, 7,10,11]
	// ]
}

func rotateImage(matrix [][]int) {
	length, last := len(matrix), len(matrix)-1

	// Debug it with i == 0 and j == 1 and then you can figure out how to rotate matrix elements.
	// leftRight < length/2; half
	// topDown < last-leftRight; another half
	for leftRight := 0; leftRight < length/2; leftRight++ {
		for topDown := leftRight; topDown < last-leftRight; topDown++ {
			tmp := matrix[leftRight][topDown]
			matrix[leftRight][topDown] = matrix[last-topDown][leftRight]
			matrix[last-topDown][leftRight] = matrix[last-leftRight][last-topDown]
			matrix[last-leftRight][last-topDown] = matrix[topDown][last-leftRight]
			matrix[topDown][last-leftRight] = tmp
		}
	}
}
