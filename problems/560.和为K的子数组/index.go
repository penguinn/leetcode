package index

func subarraySum(nums []int, k int) int {
	length := len(nums)
	if length <= 0 {
		return 0
	}

	count := 0
	now := 0

	sumMap := map[int]int{0: 1} // k为之前遍历过的和，v为个数
	for i := 0; i <= length-1; i++ {
		now += nums[i]
		if c, ok := sumMap[now-k];ok {
			count += c
		}
		sumMap[now]++
	}

	return count
}
