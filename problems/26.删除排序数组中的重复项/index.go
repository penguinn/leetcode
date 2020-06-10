package index

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 || length == 1 {
		return length
	}
	var start, end int
	end = end + 1

	for end < length {
		if nums[end] == nums[start] {
			end++
		} else {
			start++
			nums[start] = nums[end]
		}
	}
	return start + 1
}
