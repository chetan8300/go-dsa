package array

import "testing"

func TestExclusiveTime(t *testing.T) {
	tests := []struct {
		input  int
		logs   []string
		output []int
	}{
		{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}, []int{3, 4}},
	}

	for i, test := range tests {
		arr := ExclusiveTime(test.input, test.logs)
		for j, v := range arr {
			if v != test.output[j] {
				t.Fatalf("Failed test case #%d. Want %d got %d", i, test.output[j], v)
			}
		}
	}
}
