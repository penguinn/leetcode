package index

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	array := []int{0,nums[0]}
	for i:=1;i<=length-1;i++ {
		tmp := nums[i] + array[0]
		if array[1] > tmp {
			tmp = array[1]
		}
		array[0] = array[1]
		array[1] = tmp
	}
	return array[1]
}