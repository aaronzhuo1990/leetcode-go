package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	cases := []struct {
		nums []int
		ans  int
	}{

		{
			[]int{0, 1, 1, 2},
			3,
		},

		{
			[]int{0, 0, 1, 1, 2, 2, 2, 4, 5, 6},
			6,
		},
	}

	for i, c := range cases {
		result := removeDuplicatesFromSortedArray(c.nums)
		if result != c.ans {
			t.Errorf("case %d failed, expected answer %d, got result %d\n", i, c.ans, result)
		}
	}

}
