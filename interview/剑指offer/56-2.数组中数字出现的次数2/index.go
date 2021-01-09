package index

// 直接用hash表

// 位运算
func singleNumber(nums []int) int {
	res := make([]int, 32)
	for _, num := range nums {
		flag := 1
		for i := 0; i <= 31; i++ {
			if num&flag == flag {
				res[i]++
			}
			flag = flag << 1
		}
	}
	for i := 0; i <= 31; i++ {
		res[i] = res[i] % 3
	}
	result := 0
	for i := 31; i >= 0; i-- {
		result = result << 1
		result += res[i]
	}
	return result
}
