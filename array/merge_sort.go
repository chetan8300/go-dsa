package array

// MergeSort solves the problem in O(n log n) time and O(n) space.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := divide(arr)

	return merge(left, right)
}

func divide(arr []int) ([]int, []int) {
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return left, right
}

func merge(left, right []int) []int {

	output := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			output = append(output, left[i])
			i++
		} else {
			output = append(output, right[j])
			j++
		}
	}

	if i >= len(left) {
		output = append(output, right[j:]...)
	} else {
		output = append(output, left[i:]...)
	}

	return output
}
