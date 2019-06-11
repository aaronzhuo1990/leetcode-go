package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	ast := assert.New(t)

	cases := []struct {
		nums   []int
		target int
		answer []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			answer: []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 11,
			answer: []int{-1, -1},
		},
	}

	for i, c := range cases {
		result := findFirstAndLastTargetPosition(c.nums, c.target)
		ast.Equal(c.answer, result, "case %d, input: %v, target: %d\n ", i, c.nums, c.target)
	}

}
