package array

// BubbleSort solves the problem in O(n^2) time and O(1) space.
func BubbleSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}
