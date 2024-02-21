package array

import "reflect"

// import "golang.org/x/exp/maps"

func RemoveDuplicates(arr []int) []int {
	newSet := map[int]bool{}
	for i := 0; i < len(arr); i++ {
		newSet[arr[i]] = true
	}

	// intKeys := maps.Keys(newSet)
	keys := reflect.ValueOf(newSet).MapKeys()
	intKeys := make([]int, len(keys))
	for i, v := range keys {
		intKeys[i] = v.Interface().(int)
	}

	return intKeys
}
