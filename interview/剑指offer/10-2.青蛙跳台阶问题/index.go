package index

func numWays(n int) int {
	if n ==0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	nums := []int{1, 2}
	for i := 3; i <= n; i++ {
		tmp := (nums[0] + nums[1]) % (1e9 + 7)
		nums[0] = nums[1]
		nums[1] = tmp
	}

	return nums[1]
}
