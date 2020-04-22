package problems

import (
	"strconv"
)

//import "fmt"

//43.Multiply strings
//大数相乘
func multiply(num1 string, num2 string) string {
	results := make([]string, len(num2))
	resultStr := ""
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	for i := len(num2) - 1; i >= 0; i-- {
		multiplier2 := num2[i] - '0'
		carry := 0
		result := ""
		for j := len(num1) - 1; j >= 0; j-- {
			multiplier1 := num1[j] - '0'
			tmp := int(multiplier1*multiplier2) + carry
			carry = tmp / 10
			tmpResultStr := strconv.Itoa(tmp % 10)
			result = tmpResultStr + result
		}
		if carry != 0 {
			result = strconv.Itoa(carry) + result
		}

		inversionI := len(num2) - 1 - i
		for x := 0; x < inversionI; x++ {
			if result != "0" {
				result = result + "0"
			}
		}
		results[inversionI] = result
	}

	for _, tmp := range results {
		resultStr = addStrings(resultStr, tmp)
	}

	return resultStr
}

//45.Jump Game II
//1，开始的思路是用贪心找到最远的哪个染色点，然后已这个点开始做些计算，感觉这样不对
//func jump(nums []int) int {
//	count := 0
//	maxIndex := 0
//	length := len(nums)
//	for maxIndex < length-1 {
//		value := nums[maxIndex]
//		count++
//		target := maxIndex + value
//		if target >= length-1 {
//			break
//		}
//		for i := maxIndex; i <= target; i++ {
//			t := i + nums[i]
//			if t >= length-1 {
//				break
//			}
//			for ; t > maxIndex; t-- {
//				if nums[t] != 0 {
//					maxIndex = t
//					break
//				}
//			}
//		}
//		count++
//	}
//	return count
//}

//2,准备顺序执行
func jump(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	//上一步能走到的最远距离
	index := 0
	//记录最远能调到哪里
	maxIndex := 0
	//记录执行了几步
	count := 0
	for i := 0; i <= maxIndex && i < length; i++ {
		if i > index {
			count++
			index = maxIndex
		}
		if nums[i]+i > maxIndex {
			maxIndex = nums[i] + i
		}
	}
	return count
}
