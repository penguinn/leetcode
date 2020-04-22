package problems

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := [2]int{nums[0], 0}
	length := len(nums)
	for i := 1; i <= length-1; i++ {
		if i == 1 {
			if nums[1] > a[0] {
				a[1] = nums[1]
			} else {
				a[1] = a[0]
			}
			continue
		}
		dst := a[0] + nums[i]
		if dst > a[1] {
			a[0] = a[1]
			a[1] = dst
		} else {
			a[0] = a[1]
		}
	}

	return a[1]
}
