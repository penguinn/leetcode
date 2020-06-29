package index

// 使用动态规划，而没有使用递归。时间复杂度O(n)，空间复杂度O(1)
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	prepre := 0
	pre := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = (prepre + pre) % 1000000007
		prepre = pre
		pre = result
	}
	return result
}
