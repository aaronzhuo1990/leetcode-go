package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		[]int{3, 4, 5},
		[]int{0, 8, 7, 3, 4},
	}
	targets := []int{
		9,
		11,
	}
	results := [][]int{
		[]int{1, 2},
		[]int{1, 2},
	}
	caseNum := 2
	for i := 0; i < caseNum; i++ {
		if ret := twoSum(tests[i], targets[i]); ret[0] != results[i][0] && ret[1] != results[i][1] {
			t.Errorf("case %d fails: %v\n", i, ret)
		}
	}
}
