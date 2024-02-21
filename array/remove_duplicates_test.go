package array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1}},
	}

	for i, test := range tests {
		output := RemoveDuplicates(test.input)
		if !reflect.DeepEqual(MergeSort(output), MergeSort(test.output)) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}
