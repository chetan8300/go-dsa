package array

func LinearSearchBruteForce(arr []int, k int) int {
	sIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			sIndex = i
			break
		}
	}

	return sIndex
}
