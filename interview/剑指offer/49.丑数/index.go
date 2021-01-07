package index

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, 0, n)
	dp = append(dp, 1)
	i2, i3, i5 := 0, 0, 0
	for len(dp) < n {
		tmp := min(2*dp[i2], 3*dp[i3], 5*dp[i5])
		if tmp == 2*dp[i2] {
			i2++
		}
		if tmp == 3*dp[i3] {
			i3++
		}
		if tmp == 5*dp[i5] {
			i5++
		}
		dp = append(dp, tmp)
	}
	return dp[n-1]
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}
	if y <= x && y <= z {
		return y
	}
	return z
}
