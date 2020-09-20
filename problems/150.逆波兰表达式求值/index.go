package index

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	length := len(tokens)
	if length == 1 {
		num, _ := strconv.Atoi(tokens[0])
		return num
	}
	stack := make([]int, 0, len(tokens))
	var lenTmp, tmp1, tmp2 int
	for _, str := range tokens {
		if str == "+" {
			lenTmp = len(stack)
			tmp1 = stack[lenTmp-1]
			tmp2 = stack[lenTmp-2]
			stack[lenTmp-2] = tmp2 + tmp1
			stack = stack[:lenTmp-1]
		} else if str == "-" {
			lenTmp = len(stack)
			tmp1 = stack[lenTmp-1]
			tmp2 = stack[lenTmp-2]
			stack[lenTmp-2] = tmp2 - tmp1
			stack = stack[:lenTmp-1]
		} else if str == "*" {
			lenTmp = len(stack)
			tmp1 = stack[lenTmp-1]
			tmp2 = stack[lenTmp-2]
			stack[lenTmp-2] = tmp2 * tmp1
			stack = stack[:lenTmp-1]
		} else if str == "/" {
			lenTmp = len(stack)
			tmp1 = stack[lenTmp-1]
			tmp2 = stack[lenTmp-2]
			stack[lenTmp-2] = tmp2 / tmp1
			stack = stack[:lenTmp-1]
		} else {
			tmp1, _ = strconv.Atoi(str)
			stack = append(stack, tmp1)
		}
	}

	return stack[0]
}
