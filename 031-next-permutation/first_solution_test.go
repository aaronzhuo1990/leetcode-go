package next_permutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	ast := assert.New(t)
	cases := []struct {
		nums     []int
		expected []int
	}{

		{
			nums:     []int{1, 2, 7, 4, 3, 1},
			expected: []int{1, 3, 1, 2, 4, 7},
		},

		{
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
	}

	for i, c := range cases {
		nextPermutation(c.nums)
		ast.Equal(c.expected, c.nums, "case %d, input %v", i, c.nums)
	}

}
