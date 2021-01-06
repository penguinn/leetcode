package index

func exchange(nums []int) []int {
	length := len(nums)
	j := 0
	for i:=0;i<=length-1;i++ {
		if nums[i]%2==1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
