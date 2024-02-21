package array

import "sort"

func MissingNumberBruteForce(nums []int) int {
	sort.Ints(nums)
	var missingNumber int

	for i := 0; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != i {
			missingNumber = i
			break
		}
	}

	return missingNumber
}

func MissingNumberBest(nums []int) int {
	n := len(nums)

	totalSum := (n * (n + 1)) / 2
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return totalSum - sum

	// n := len(nums)
	// xor1, xor2 := 0, 0

	// for i := 0; i < n; i++ {
	// 	xor2 = xor2 ^ nums[i]
	// 	xor1 = xor1 ^ (i + 1)
	// }

	// xor1 = xor1 ^ (n + 1)
	// return xor1 ^ xor2

	// x := len(nums)
	// for i, val := range nums {
	// 	x = x ^ val
	// 	x = x ^ i
	// }
	// return x
}
