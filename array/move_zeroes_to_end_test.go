package array

import (
	"reflect"
	"testing"
)

func TestMoveZeroesToEnd(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{1, 0, 0}, []int{1, 0, 0}},
		{[]int{0, 1, 0}, []int{1, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 0, 2}, []int{1, 2, 0}},
		{[]int{0, 1, 2}, []int{1, 2, 0}},
		{[]int{1, 2, 0}, []int{1, 2, 0}},
		{[]int{0, 0, 1, 0, 2, 0, 3, 0}, []int{1, 2, 3, 0, 0, 0, 0, 0}},
	}

	for i, test := range tests {
		output := MoveZeroesToEnd(test.input)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}

func TestMoveZeroesToEndOptimal(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{1, 0, 0}, []int{1, 0, 0}},
		{[]int{0, 1, 0}, []int{1, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 0, 2}, []int{1, 2, 0}},
		{[]int{0, 1, 2}, []int{1, 2, 0}},
		{[]int{1, 2, 0}, []int{1, 2, 0}},
		{[]int{0, 0, 1, 0, 2, 0, 3, 0}, []int{1, 2, 3, 0, 0, 0, 0, 0}},
	}

	for i, test := range tests {
		output := MoveZeroesToEndOptimal(test.input)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}
