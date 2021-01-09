package index

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	queueLength := 0
	result := []int{}
	length := len(nums)
	for i := 0; i <= length-1; i++ {
		for queueLength != 0 && nums[i] > queue[queueLength-1] {
			queueLength--
		}
		queue = queue[:queueLength]
		queue = append(queue, nums[i])
		queueLength++
		if i > k && nums[i-k] == queue[0] {
			queue = queue[1:]
			queueLength--
		}
		if i >= k-1 {
			result = append(result, queue[0])
		}
	}
	return result
}
