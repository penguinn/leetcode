package index

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	fmt.Println(num)
	result := 0
	for num > 0 {
		if num&1 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}
