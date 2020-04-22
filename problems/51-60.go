package problems

import (
	"math"
)

//53. Maximum Subarray
//1.开头的数不能是负数，相加的和如果是负数了，开始指针就跳到结束指针后一个数重新开始
//2.吸取，把相加的结果用一个数组保存
func MaxSubArray(nums []int) int {
	tmp := make([]int, len(nums))
	tmp[0] = nums[0]
	sum := tmp[0]
	for i := 1; i < len(nums); i++ {
		if tmp[i-1] < 0 {
			tmp[i] = nums[i]
		} else {
			tmp[i] = tmp[i-1] + nums[i]
		}
		if sum < tmp[i] {
			sum = tmp[i]
		}
	}
	return sum
}

//55 jump Game
//遍历数组前行，把能染色的都染色了
//1, 遍历整个切片染色，会出现超时
//func canJump(nums []int) bool {
//	signs := map[int]bool{0: true}
//	for i, num := range nums {
//		if signs[i] {
//			for j := 1; j <= num; j++ {
//				signs[i+j] = true
//			}
//		}
//	}
//
//	if _, ok := signs[len(nums)-1]; ok {
//		return true
//	} else {
//		return false
//	}
//}
//2，标识出染色的最大值，这其实也是贪心算法的一种变种
func canJump(nums []int) bool {
	max := 0
	for i, num := range nums {
		if i > max {
			return false
		}
		max = int(math.Max(float64(max), float64(i+num)))
	}
	return true
}
