package index

import (
	"strconv"
	"strings"
)

// 时间复杂度O(n), 空间复杂度O(n)
func calculate(s string) int {
	slices := strings.Split(s, "")
	result := 0
	current := 0
	sign := 1
	queue := []string{}

	for _, element := range slices {
		switch element {
		case " ":
			continue
		case "+":
			result += current * sign
			sign = 1
			current = 0
		case "-":
			result += current * sign
			sign = -1
			current = 0
		case "(":
			queue = append(queue, strconv.Itoa(result))
			queue = append(queue, strconv.Itoa(sign))
			result = 0
			current = 0
			sign = 1
		case ")":
			result += current * sign
			current = 0
			tmpSign, _ := strconv.Atoi(queue[len(queue)-1])
			queue = queue[:len(queue)-1]
			tmpResult, _ := strconv.Atoi(queue[len(queue)-1])
			queue = queue[:len(queue)-1]
			result = tmpResult + tmpSign*result
		default:
			tmpNum, _ := strconv.Atoi(element)
			current = tmpNum + 10*current
		}
	}

	return result + current*sign
}
