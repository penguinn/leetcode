package index

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	result := [][]int{}
	var target int
	for i := 0; i <= length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			continue
		}
		target = 0 - nums[i]
		j := i + 1
		k := length - 1
		for j < k {
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k != length-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[j]+nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
