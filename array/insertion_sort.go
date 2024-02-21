package array

// InsertionSort solves the problem in O(n^2) time and O(1) space.
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for i-1 >= 0 && arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}

	return arr
}
