package array

import (
	"testing"
)

func TestLinearSearchBruteForce(t *testing.T) {
	tests := []struct {
		arr []int
		k   int
		out int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 8},
	}
	for _, test := range tests {
		out := LinearSearchBruteForce(test.arr, test.k)
		if out != test.out {
			t.Errorf("Test Failed: expected %v, got %v", test.out, out)
		}
	}
}
