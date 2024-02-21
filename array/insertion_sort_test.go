package array

import (
	"reflect"
	"testing"
)

// SelectionSort solves the problem in O(n^2) time and O(1) space.
func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input  []int
		sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for i, test := range tests {
		InsertionSort(test.input)
		if !reflect.DeepEqual(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.sorted, test.input)
		}
	}
}
