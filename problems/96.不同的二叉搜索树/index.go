package index

// 使用递归的方式
var tmpMap = map[int]int{}

func numTrees(n int) int {
	tmpMap = map[int]int{}
	return numTree(n)
}

func numTree(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if sum, ok := tmpMap[n]; ok {
		return sum
	}
	sum := 0
	for i := 0; i <= n-1; i++ {
		sum += numTree(i) * numTree(n-i-1)
	}
	tmpMap[n] = sum
	return sum
}

// 使用动态规划
func numTrees1(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}
