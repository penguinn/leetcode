package index

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []int{}
	times := 1
	result := ""
	tmpTimes := 0
	for _, c := range strings.Split(s, "") {
		if c == "[" {
			stack = append(stack, tmpTimes)
			times *= tmpTimes
			tmpTimes = 0
		} else if c == "]" {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmpTimes /= tmp
		} else {
			inst, err := strconv.Atoi(c)
			if err == nil {
				tmpTimes = 10*tmpTimes + inst
			} else {
				for i := 0; i < times; i++ {
					result += c
				}
			}
		}
	}

	return result
}
