package index

import (
	"fmt"
)

func singleNumbers(nums []int) []int {
	result := 0
	for _, num := range nums {
		result = result ^ num
	}
	mask := 1
	for mask&result == 0 {
		mask = mask << 1
	}

	fmt.Println(result, mask)

	a := 0
	b := 0
	for _, num := range nums {
		if num&mask == 0 {
			fmt.Println(num)
			a = a ^ num
		} else {
			b = b ^ num
		}
	}

	return []int{a, b}
}
