package problems
/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
func rob11(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}
	sum := make([]int, length)

	for i, num := range nums {
		if i == 0 {
			sum[i] = num
		} else if i == 1 {
			if num > sum[0] {
				sum[i] = num
			} else {
				sum[i] = sum[0]
			}
		} else {
			if sum[i-1] > sum[i-2]+num {
				sum[i] = sum[i-1]
			} else {
				sum[i] = sum[i-2] + num
			}
		}
	}
	return sum[length-1]
}

