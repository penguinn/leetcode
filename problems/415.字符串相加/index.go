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

func addStrings1(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	length1 := len(num1)
	length2 := len(num2)

	result := ""
	i, j := length1-1, length2-1
	carray := 0
	for i >= 0 && j >= 0 {
		add1 := num1[i] - '0'
		add2 := num2[j] - '0'
		sum := int(add1) + int(add2) + carray
		carray = sum / 10
		result = strconv.Itoa(sum % 10) + result
		i--
		j--
	}

	for i >= 0 {
		add := num1[i] - '0'
		sum := int(add) + carray
		carray = sum / 10
		result = strconv.Itoa(sum % 10) + result
		i--
	}

	for j >= 0 {
		add := num2[j] - '0'
		sum := int(add) + carray
		carray = sum / 10
		result = strconv.Itoa(sum % 10) + result
		j--
	}

	if carray == 1 {
		result = "1" + result
	}

	return result
}
