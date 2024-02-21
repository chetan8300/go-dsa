package array

import "slices"

func largestElementInArrayBruteForce(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	slices.Sort(arr)
	return arr[len(arr)-1]
}

func largestElementInArrayRecursive(arr []int) int {
	if len(arr) > 0 {
		max := arr[0]

		for i := 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
			}
		}

		return max
	}

	return 0
}

func secondLargestElementInArrayBruteForce(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	slices.Sort(arr)

	secondMax := arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < secondMax {
			secondMax = arr[i]
			break
		}
	}

	return secondMax
}

func secondLargestElementInArrayBestApproach(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	max := largestElementInArrayRecursive(arr)

	secondLarge := arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if secondLarge < arr[i] && arr[i] != max {
			secondLarge = arr[i]
		}
	}

	return secondLarge
}

func SecondLargestElementInArrayOptimalApproach(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	largest := arr[0]
	secondLargest := arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if largest < arr[i] {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] < largest && arr[i] > secondLargest {
			secondLargest = arr[i]
		}
	}

	return secondLargest
}
