package search_rotated_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	ast := assert.New(t)

	cases := []struct {
		nums   []int
		target int
		answer int
	}{
		//{
		//	nums:   []int{4,5,6,7,0,1,2},
		//	target: 6,
		//	answer: 2,
		//},
		//{
		//	nums:   []int{4,5,6,7,0,1,2},
		//	target: 1,
		//	answer: 5,
		//},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 8,
			answer: -1,
		},
	}

	for i, c := range cases {
		result := searchInRotatedSortedArray(c.nums, c.target)
		ast.Equal(c.answer, result, "case %d, input: %v, target: %d\n ", i, c.nums, c.target)
	}

}
