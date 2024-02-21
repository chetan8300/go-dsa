package array

func MaximumConsecutiveOnes(nums []int) int {
	prevLongest := 0
	currLongest := 0

	for _, v := range nums {
		if v == 1 {
			currLongest++
		} else {
			prevLongest = max(prevLongest, currLongest)
			currLongest = 0
		}
	}

	return max(prevLongest, currLongest)
}
