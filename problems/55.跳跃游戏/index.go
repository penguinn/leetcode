package index

// 利用贪心
func canJump(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}
	maxIndex := 0
	for i, num := range nums {
		if i > maxIndex {
			break
		} else if i+num > maxIndex {
			maxIndex = i + num
		}
	}
	if maxIndex >= length-1 {
		return true
	} else {
		return false
	}
}
