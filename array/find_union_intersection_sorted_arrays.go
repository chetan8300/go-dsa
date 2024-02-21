package array

func FindUnionInSortedArrays(arr1 []int, arr2 []int) []int {
	setOfUnion := map[int]bool{}

	for _, v := range arr1 {
		setOfUnion[v] = true
	}

	for _, v := range arr2 {
		setOfUnion[v] = true
	}

	union := []int{}

	for number := range setOfUnion {
		union = append(union, number)
	}

	return union
}

func FindIntersectionInSortedArrays(arr1 []int, arr2 []int) []int {
	intersection := []int{}
	n1, n2 := len(arr1), len(arr2)
	visited := make(map[int]int, n2)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if arr1[i] == arr2[j] && visited[j] != 1 {
				visited[j] = 1
				intersection = append(intersection, arr1[i])
				break
			}

			if arr2[j] > arr1[i] {
				break
			}
		}
	}

	return intersection
}

func FindUnionInSortedArraysOptimalApproach(arr1 []int, arr2 []int) []int {
	union := []int{}
	i, j := 0, 0
	n1, n2 := len(arr1), len(arr2)

	for i < n1 && j < n2 {
		if arr1[i] <= arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}
	}

	for j < n2 {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}
	for i < n1 {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	return union
}

func FindIntersectionInSortedArraysOptimalApproach(nums1 []int, nums2 []int) []int {
	intersection := []int{}
	i, j := 0, 0
	n1, n2 := len(nums1), len(nums2)

	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			intersection = append(intersection, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return intersection
}
