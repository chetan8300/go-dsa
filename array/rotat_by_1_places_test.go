package array

import (
	"reflect"
	"testing"
)

func TestRotateBy1Place(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 2, 1}, []int{2, 1, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{4, 3, 2, 1, 5}},
	}

	for i, test := range tests {
		output := RotateBy1Place(test.input, 1)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}

func TestRotateByKPlacesBruteForce(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{}, 1, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{2, 1}, 1, []int{1, 2}},
		{[]int{3, 2, 1}, 1, []int{2, 1, 3}},
		{[]int{5, 4, 3, 2, 1}, 1, []int{4, 3, 2, 1, 5}},
		{[]int{5, 4, 3, 2, 1}, 2, []int{3, 2, 1, 5, 4}},
		{[]int{5, 4, 3, 2, 1}, 3, []int{2, 1, 5, 4, 3}},
		{[]int{5, 4, 3, 2, 1}, 4, []int{1, 5, 4, 3, 2}},
		{[]int{5, 4, 3, 2, 1}, 5, []int{5, 4, 3, 2, 1}},
	}

	for i, test := range tests {
		output := RotateByKPlacesBruteForce(test.input, test.k)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}

func TestRotateByKPlacesBetter(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{}, 1, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{2, 1}, 1, []int{1, 2}},
		{[]int{3, 2, 1}, 1, []int{2, 1, 3}},
		{[]int{5, 4, 3, 2, 1}, 1, []int{4, 3, 2, 1, 5}},
		{[]int{5, 4, 3, 2, 1}, 2, []int{3, 2, 1, 5, 4}},
		{[]int{5, 4, 3, 2, 1}, 3, []int{2, 1, 5, 4, 3}},
		{[]int{5, 4, 3, 2, 1}, 4, []int{1, 5, 4, 3, 2}},
		{[]int{5, 4, 3, 2, 1}, 5, []int{5, 4, 3, 2, 1}},
	}

	for i, test := range tests {
		output := RotateByKPlacesBetter(test.input, test.k)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}

func TestRotateByKPlacesOptimal(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{}, 1, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{2, 1}, 1, []int{1, 2}},
		{[]int{3, 2, 1}, 1, []int{2, 1, 3}},
		{[]int{5, 4, 3, 2, 1}, 1, []int{4, 3, 2, 1, 5}},
		{[]int{5, 4, 3, 2, 1}, 2, []int{3, 2, 1, 5, 4}},
		{[]int{5, 4, 3, 2, 1}, 3, []int{2, 1, 5, 4, 3}},
		{[]int{5, 4, 3, 2, 1}, 4, []int{1, 5, 4, 3, 2}},
		{[]int{5, 4, 3, 2, 1}, 5, []int{5, 4, 3, 2, 1}},
	}

	for i, test := range tests {
		output := RotateByKPlacesOptimal(test.input, test.k)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.output, output)
		}
	}
}
