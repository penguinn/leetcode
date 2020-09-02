package index

func majorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	// 一个数的数量
	resultMap := map[int]int{}
	var current int
	var ok bool
	for _, num := range nums {
		current, ok = resultMap[num]
		if !ok {
			resultMap[num] = 1
		} else {
			if current + 1 > length/2 {
				return num
			} else {
				resultMap[num] = current+1
			}
		}
	}

	return -1
}
