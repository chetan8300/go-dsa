package array

// SelectionSort solves the problem in O(n^2) time and O(1) space.
func SelectionSort(input []int) {
	for i := 0; i < len(input); i++ {
		minIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}
}
