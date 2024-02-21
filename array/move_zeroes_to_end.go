package array

func MoveZeroesToEnd(arr []int) []int {
	temp := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp = append(temp, arr[i])
		}
	}

	zeroes := len(arr) - len(temp)
	for j := 0; j < zeroes; j++ {
		temp = append(temp, 0)
	}

	return temp
}

func MoveZeroesToEndOptimal(arr []int) []int {
	j := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}

	if j != -1 {
		for i := j + 1; i < len(arr); i++ {
			if arr[i] != 0 {
				arr[i], arr[j] = arr[j], arr[i]
				j++
			}
		}
	}

	return arr
}
