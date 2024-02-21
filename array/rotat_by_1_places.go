package array

func RotateBy1Place(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}

	temp := arr[0]

	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}

	arr[len(arr)-1] = temp

	return arr
}

func RotateByKPlacesBruteForce(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}

	k = k % len(arr)

	for i := 0; i < k; i++ {
		temp := arr[0]

		for j := 1; j < len(arr); j++ {
			arr[j-1] = arr[j]
		}

		arr[len(arr)-1] = temp
	}

	return arr
}

func RotateByKPlacesBetter(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}

	k = k % len(arr)

	temp := []int{}

	for i := 0; i < k; i++ {
		temp = append(temp, arr[i])
	}

	arr = arr[k:]

	for i := 0; i < len(temp); i++ {
		arr = append(arr, temp[i])
	}

	return arr
}

func PartialReverse(arr []int, start, end int) {
	for start <= end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func RotateByKPlacesOptimal(arr []int, k int) []int {
	if len(arr) <= 1 {
		return arr
	}

	k = k % len(arr)

	PartialReverse(arr, 0, k-1)
	PartialReverse(arr, k, len(arr)-1)
	PartialReverse(arr, 0, len(arr)-1)

	return arr
}
