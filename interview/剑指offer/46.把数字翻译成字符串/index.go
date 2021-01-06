package index

import (
	"strconv"
)

// 递归思想
func translateNum1(num int) int {
	if num <= 0 {
		return 1
	} else if num%100 <= 25 && num%100 >= 10 {
		return translateNum1(num/100) + translateNum1(num/10)
	} else {
		return translateNum1(num / 10)
	}
}

// 动态规划思想
func translateNum(num int) int {
	if num < 0 {
		return 0
	}
	if num == 0 {
		return 1
	}
	str := strconv.Itoa(num)
	length := len(str)
	result := []int{1, 1}
	tmp := 0
	for i := 0; i <= length-2; i++ {
		if str[i:i+2] >= "10" && str[i:i+2] <= "25" {
			tmp = result[0] + result[1]
			result[0] = result[1]
			result[1] = tmp
		} else {
			result[0] = result[1]
		}
	}

	return result[1]
}
