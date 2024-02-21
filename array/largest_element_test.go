package array

import "testing"

func TestLargestElementInArray(t *testing.T) {
	tests := []struct {
		input         []int
		largest       int
		secondLargest int
	}{
		{[]int{}, 0, 0},
		{[]int{1}, 1, 1},
		{[]int{2, 1}, 2, 1},
		{[]int{3, 2, 1}, 3, 2},
		{[]int{5, 4, 3, 2, 1}, 5, 4},
	}

	for i, test := range tests {
		largest := largestElementInArrayBruteForce(test.input)
		if largest != test.largest {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.largest, largest)
		}
	}

	for i, test := range tests {
		largest := largestElementInArrayRecursive(test.input)
		if largest != test.largest {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.largest, largest)
		}
	}

	for i, test := range tests {
		secondLargest := secondLargestElementInArrayBruteForce(test.input)
		if secondLargest != test.secondLargest {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.secondLargest, secondLargest)
		}
	}

	for i, test := range tests {
		secondLargest := secondLargestElementInArrayBestApproach(test.input)
		if secondLargest != test.secondLargest {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.secondLargest, secondLargest)
		}
	}

	// for i, test := range tests {
	// 	secondLargest := secondLargestElementInArrayOptimalApproach(test.input)
	// 	if secondLargest != test.secondLargest {
	// 		t.Fatalf("Failed test case #%d. Want %d got %d", i, test.secondLargest, secondLargest)
	// 	}
	// }
}
