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
		target = 0 - nums[i]
		for j := i + 1; j <= length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := length - 1
			for j != k {
				if k != length-1 && nums[k] == nums[k+1] {
					continue
				}
				if nums[j]+nums[k] == target {
					result = append(result, []int{i, j, k})
					break
				} else if nums[j]+nums[k] > target {
					k--
					continue
				} else {
					break
				}
			}
		}
	}
	return result
}
