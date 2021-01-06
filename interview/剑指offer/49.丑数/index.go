package index

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	nums := make([]int, 0, n)
	nums = append(nums, 1)
	i2, i3, i5 := 1, 1, 1
	for len(nums) < n {
		tmp := min(2*i2, 3*i3, 5*i5)
		if tmp == 2*i2 {
			i2++
		} else if tmp == 3*i3 {
			i3++
		} else {
			i5++
		}
		nums = append(nums, tmp)
	}
	return nums[n-1]
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < x && y < z {
		return y
	}
	return z
}
