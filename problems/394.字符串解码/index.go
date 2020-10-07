package index

import (
	"strconv"
	"strings"
)

type stackStr struct {
	Str   string // "["前的字符串
	Times int    // "["前的次数
}

func decodeString(s string) string {
	result := ""
	stack := []stackStr{}
	times := 0
	str := ""
	for _, c := range strings.Split(s, "") {
		if c == "[" {
			stack = append(stack, stackStr{Str: str, Times: times})
			times = 0
			str = ""
		} else if c == "]" {
			tmpElement := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			str = strings.Repeat(str, tmpElement.Times)
			str = tmpElement.Str + str
			if len(stack) == 0 {
				result += str
				str = ""
			}
		} else {
			inst, err := strconv.Atoi(c)
			if err == nil {
				times = 10*times + inst
			} else {
				str += c
			}
		}
	}

	// 最后的字符串
	result += str

	return result
}
