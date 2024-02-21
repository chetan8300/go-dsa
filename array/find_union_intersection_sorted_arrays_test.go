package array

import (
	"testing"
)

func TestFindUnionInSortedArray(t *testing.T) {
	arr1 := []int{1, 2, 2, 3, 3, 4, 5, 6}
	arr2 := []int{2, 3, 3, 5, 6, 7}

	union := FindUnionInSortedArrays(arr1, arr2)

	if len(union) != 7 {
		t.Errorf("Expected 7, got %d", len(union))
	}
}

func TestFindIntersectionInSortedArray(t *testing.T) {
	arr1 := []int{1, 2, 2, 3, 3, 4, 5, 6}
	arr2 := []int{2, 3, 3, 5, 6, 7}

	intersection := FindIntersectionInSortedArrays(arr1, arr2)

	if len(intersection) != 5 {
		t.Errorf("Expected 4, got %d", len(intersection))
	}
}

func TestFindUnionInSortedArrayOptimal(t *testing.T) {
	arr1 := []int{1, 3, 4, 5, 7}
	arr2 := []int{2, 3, 5, 6}

	union := FindUnionInSortedArraysOptimalApproach(arr1, arr2)

	if len(union) != 7 {
		t.Errorf("Expected 7, got %d", len(union))
	}
}

func TestFindIntersectionInSortedArrayOptimal(t *testing.T) {
	arr1 := []int{1, 2, 2, 3, 3, 4, 5, 6}
	arr2 := []int{2, 3, 3, 5, 6, 7}

	intersection := FindIntersectionInSortedArraysOptimalApproach(arr1, arr2)

	if len(intersection) != 5 {
		t.Errorf("Expected 4, got %d", len(intersection))
	}
}
