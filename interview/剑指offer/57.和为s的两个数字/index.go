package index

func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	left := 0
	right := length - 1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
