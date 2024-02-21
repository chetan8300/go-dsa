package array

// Time: O(n log n), Space: O(1)
func QuickSort(arr []int) []int {
	qs(arr, 0, len(arr)-1)

	return arr
}

func qs(arr []int, low, high int) {
	if low < high {
		pIndex := findPartitionAndSwap(arr, low, high)
		qs(arr, low, pIndex-1)
		qs(arr, pIndex+1, high)
	}
}

func findPartitionAndSwap(arr []int, low, high int) int {
	pivot := arr[low] // 2
	i, j := low, high // 0 7

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]

	return j
}
