package problems

import (
	"strconv"
)

//415. AddGithub Strings
//大数相加

func addStrings(num1 string, num2 string) string {
	maxNum := ""
	carry := 0
	result := ""
	if len(num1) > len(num2) {
		maxNum = num1
	} else {
		maxNum = num2
	}
	for i := 0; i < len(maxNum); i++ {
		var addend1 int
		if i1 := len(num1) - 1 - i; i1 >= 0 {
			char1 := num1[i1]
			addend1 = int(char1 - '0')
		}
		var addend2 int
		if i2 := len(num2) - 1 - i; i2 >= 0 {
			char2 := num2[i2]
			addend2 = int(char2 - '0')
		}

		tmp := carry + addend1 + addend2
		if tmp >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		tmpResultStr := strconv.Itoa(tmp % 10)
		result = tmpResultStr + result
	}

	if carry == 1 {
		result = "1" + result
	}
	return result
}
