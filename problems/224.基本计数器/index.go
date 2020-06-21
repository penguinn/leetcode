package index

import (
	"strconv"
	"strings"
)

// 时间复杂度O(n), 空间复杂度O(n)
func calculate(s string) int {
	slices := strings.Split(s, "")
	intputQueue := []string{}
	var result int
	var operate int
	sign := 1

	for _, str := range slices {
		if str == " " {
			continue
		} else if str == "+" {
			result += operate * sign
			operate = 0
			sign = 1
		} else if str == "-" {
			result += operate * sign
			operate = 0
			sign = -1
		} else if str == "(" {
			intputQueue = append(intputQueue, strconv.Itoa(result))
			intputQueue = append(intputQueue, strconv.Itoa(sign))
			result = 0
			operate = 0
			sign = 1
		} else if str == ")" {
			result += operate * sign
			operate = 0
			tmpSign, _ := strconv.Atoi(intputQueue[len(intputQueue)-1])
			intputQueue = intputQueue[:len(intputQueue)-1]
			tmpResult, _ := strconv.Atoi(intputQueue[len(intputQueue)-1])
			intputQueue = intputQueue[:len(intputQueue)-1]
			result = tmpResult + tmpSign*result
		} else {
			tmpNum, _ := strconv.Atoi(str)
			operate = 10*operate + tmpNum
		}
	}

	return result + sign*operate
}
