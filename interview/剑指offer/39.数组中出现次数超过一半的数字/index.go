package index


// 第一种方法直接使用哈希表

// 第二种投票算法
func majorityElement(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	candidate := nums[0]
	count := 1
	for i:=1;i<=length-1;i++ {
		if count == 0 {
			candidate = nums[i]
			count++
		} else {
			if candidate == nums[i] {
				count++
			} else {
				count--
			}
		}
	}

	return candidate
}
