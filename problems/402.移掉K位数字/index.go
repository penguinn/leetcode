package index

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	result := []byte{}

	for i, _ := range num {
		for k>0 && len(result) > 0 && num[i] < result[len(result)-1] {
			result = result[:len(result)-1]
			k--
		}
		result = append(result, num[i])
	}

	tmp := strings.TrimLeft(string(result[:len(result)-k]), "0")
	if tmp == "" {
		return "0"
	}
	return tmp
}