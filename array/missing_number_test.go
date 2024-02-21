package array

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{0, 1, 2, 3, 4, 6, 7, 8}, 5},
		{[]int{0, 1, 2, 3, 4, 5, 7, 8, 9}, 6},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 10}, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10}, 8},
	}

	for i, test := range tests {
		missingNumber := MissingNumberBruteForce(test.input)
		if missingNumber != test.output {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.output, missingNumber)
		}
	}

	for i, test := range tests {
		missingNumber := MissingNumberBest(test.input)
		if missingNumber != test.output {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.output, missingNumber)
		}
	}
}
