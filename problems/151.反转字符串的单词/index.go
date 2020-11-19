package index

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	slices := strings.Split(s, " ")
	result := []string{}
	for i := len(slices) - 1; i >= 0; i-- {
		if strings.TrimSpace(slices[i]) != "" {
			result = append(result, slices[i])
		}
	}

	return strings.Join(result, " ")
}
