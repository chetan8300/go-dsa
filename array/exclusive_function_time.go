package array

import (
	"strconv"
	"strings"
)

func ExclusiveTime(n int, logs []string) []int {
	arr := make([]int, n)
	stack := []int{}
	prevTimeStamp := 0

	for i := 0; i < len(logs); i++ {
		splitString := strings.Split(logs[i], ":")

		fId, ferr := strconv.Atoi(splitString[0])
		eType := splitString[1]
		timestamp, terr := strconv.Atoi(splitString[2])

		if terr == nil && ferr == nil {
			if eType == "start" {
				if len(arr) > 0 {
					arr[len(arr)-1] = timestamp - prevTimeStamp
				}

				stack = append(stack, fId)
				prevTimeStamp = int(timestamp)
			} else {
				lastElem := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				arr[lastElem] += timestamp - prevTimeStamp + 1
				prevTimeStamp = timestamp + 1
			}
		}
	}

	return arr
}
