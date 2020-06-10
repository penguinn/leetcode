package index

// 注意临界点
func nextPermutation(nums []int) {
	length := len(nums)
	if length == 0 || length == 1 {
		return
	}

	index := -1
	for i := length - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i - 1
			break
		}
	}

	if index != -1 {
		for i := length - 1; i > index; i-- {
			if nums[i] > nums[index] {
				nums[i], nums[index] = nums[index], nums[i]
				break
			}
		}
	}
	for i := 0; i < (length-index)/2; i++ {
		nums[index+1+i], nums[length-1-i] = nums[length-1-i], nums[index+1+i]
	}
	return
}
