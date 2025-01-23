package minmax

import (
	"testing"
)

func TestMinmax(t *testing.T) {
	tests := []struct {
		list        []int
		expectedMin int
		expectedMax int
	}{
		{[]int{3, 5, 7, 2, 8, -1, 4, 10, 12}, -1, 12},
		{[]int{1}, 1, 1},
		{[]int{-5, -9, -3, -4}, -9, -3},
		{[]int{100, 200, 300, 400}, 100, 400},
		{[]int{0, 0, 0, 0}, 0, 0},
	}

	for _, test := range tests {
		min, max := minmax(test.list)
		if min != test.expectedMin || max != test.expectedMax {
			t.Errorf("For list %v, expected min %d and max %d, but got min %d and max %d", test.list, test.expectedMin, test.expectedMax, min, max)
		}
	}
}
