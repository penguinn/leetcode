package index

// 利用贪心
func canJump(nums []int) bool {
	length := len(nums)
	var max int
	for i:=0;i<=max;i++ {
		if i+nums[i] > max {
			max = i+nums[i]
		}
		if max >= length-1 {
			return true
		}
	}
	return false
}
