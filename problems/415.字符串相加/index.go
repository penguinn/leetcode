package index

import (
	"strconv"
)

// 记得减去0字符
func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	if len1 == 0 {
		return num2
	}
	len2 := len(num2)
	if len2 == 0 {
		return num1
	}
	flag := 0
	result := ""
	i := 1
	for ; i <= len1 && i <= len2; i++ {
		a := num1[len1-i] - '0'
		b := num2[len2-i] - '0'
		val := int(a) + int(b) + flag
		flag = val / 10
		result = strconv.Itoa(val%10) + result
	}

	for ; i <= len1; i++ {
		a := num1[len1-i] - '0'
		val := int(a) + flag
		flag = val / 10
		result = strconv.Itoa(val%10) + result
	}

	for ; i <= len2; i++ {
		b := num2[len2-i] - '0'
		val := int(b) + flag
		flag = val / 10
		result = strconv.Itoa(val%10) + result
	}

	if flag != 0 {
		result = "1" + result
	}

	return result
}
