package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	cases := []struct {
		heights []int
		ans     int
	}{

		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},

		{
			[]int{1, 2, 3, 4, 3, 5, 2},
			10,
		},
	}

	for i, c := range cases {
		result := findMaxArea(c.heights)
		if result != c.ans {
			t.Errorf("case %d failed, expected answer %d, got result %d\n", i, c.ans, result)
		}
	}

}
