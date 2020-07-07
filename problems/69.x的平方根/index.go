package index

import (
	"fmt"
	"strconv"
)

func mySqrt(x int) int {
	start := 0
	end := x - 1
	for start <= end {
		mid := start + ((end - start) >> 1)
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if (mid+1)*(mid+1) == x {
			return mid + 1
		} else if (mid+1)*(mid+1) < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return 0
}

// 获取四舍五入的小数点后6位开方
func mySqrt6(x float64) float64 {
	start := 0.000000
	end := x - 1
	for start <= end {
		mid := Decimal7(start + (end-start)/2)
		if mid*mid <= x && (mid+0.0000001)*(mid+0.0000001) > x {
			return Decimal6(mid)
		} else if (mid+0.0000001)*(mid+0.0000001) == x {
			return Decimal6(mid + 0.0000001)
		} else if (mid+0.0000001)*(mid+0.0000001) < x {
			start = mid + 0.0000001
		} else {
			end = mid - 0.0000001
		}
	}
	return 0
}

func Decimal6(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.6f", value), 64)
	return value
}

func Decimal7(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.7f", value), 64)
	return value
}
