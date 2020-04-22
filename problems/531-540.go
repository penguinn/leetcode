package problems

import (
	"strconv"
	"strings"
)

//537. Complex Number Multiplication
//第一想法还是硬解，通过for循环来找到+和i的下标
//第二想法在考虑利用split
//这题没啥难度，也没太大意义
func ComplexNumberMultiply(a string, b string) string {

	aArray1 := strings.Split(a, "+")
	aReal, _ := strconv.Atoi(aArray1[0])
	aImag, _ := strconv.Atoi(strings.Split(aArray1[1], "i")[0])

	bArray1 := strings.Split(b, "+")
	bReal, _ := strconv.Atoi(bArray1[0])
	bImag, _ := strconv.Atoi(strings.Split(bArray1[1], "i")[0])

	cReal := aReal*bReal - aImag*bImag
	cImag := aReal*bImag + aImag*bReal

	c := strconv.Itoa(cReal) + "+" + strconv.Itoa(cImag) + "i"

	return c
}
