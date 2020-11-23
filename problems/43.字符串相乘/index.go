package index

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	length1 := len(num1)
	length2 := len(num2)

	array := make([]int, length1+length2-1)
	for i := 0; i <= length1-1; i++ {
		for j := 0; j <= length2-1; j++ {
			multi1 := num1[i] - '0'
			multi2 := num2[j] - '0'
			array[i+j] += int(multi1) * int(multi2)
		}
	}
	carry := 0
	for i := length1 + length2 - 2; i >= 0; i-- {
		tmp := array[i] + carry
		carry = tmp / 10
		array[i] = tmp % 10
	}
	if carry != 0 {
		array = append([]int{carry}, array...)
	}

	result := ""
	for i := 0; i <= len(array)-1; i++ {
		result = result + strconv.Itoa(array[i])
	}

	return result
}
