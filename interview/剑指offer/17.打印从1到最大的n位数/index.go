package index

import (
	"strconv"
)

func printNumbers(n int) []int {
	if n <=0 {
		return []int{}
	}
	num := ""
	for i:=0;i<=n-1;i++ {
		num+="9"
	}
	maxNum, _ := strconv.Atoi(num)
	result := make([]int, 0, maxNum)
	for i:=1;i<=maxNum;i++ {
		result = append(result, i)
	}
	return result
}
