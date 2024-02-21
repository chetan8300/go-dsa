package array

func CheckIfSorted(arr []int) bool {
	sorted := true

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] < arr[i-1] {
			sorted = false
			break
		}
	}

	return sorted
}
