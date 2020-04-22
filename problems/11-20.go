package problems

import (
	"fmt"
	"math"
	"sort"
)

//15.3Sum
//想到利用twoSum, 但是这个的时间复杂度是0(n*n)
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := [][]int{}
	for i, ch := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		tmpResult := twoSum(nums[i+1:], 0-ch)
		if len(tmpResult) != 0 {
			for _, oneResult := range tmpResult {
				oneResult = append(oneResult, ch)
				result = append(result, oneResult)
			}

		}
	}
	return result
}

func twoSum(nums []int, target int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}
	numMap := map[int]int{}
	result := [][]int{}
	for i, ch := range nums {
		if i != 0 && nums[i] == nums[i-1] && len(result) != 0 {
			continue
		}
		if _, ok := numMap[target-ch]; ok {
			oneResult := make([]int, 0, 2)
			oneResult = append(oneResult, ch, target-ch)
			result = append(result, oneResult)
		} else {
			numMap[ch] = 1
		}
	}
	fmt.Println(result)
	return result
}

//16.3Sum Closest
//这道题没法用hashMap了。如果用暴力解法，时间复杂度为O(n*n*n)
//可以用双指针的方式，上题也可以这样使用，时间复杂度都为O(n*n)
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	globalDiff := math.MaxInt32
	sort.Ints(nums)
	for i, num := range nums {
		tmpTarget := target - num
		start := i + 1
		end := length - 1
		for start < end {
			diff := nums[start] + nums[end] - tmpTarget
			if int(math.Abs(float64(diff))) < int(math.Abs(float64(globalDiff))) {
				globalDiff = diff
			}
			if diff > 0 {
				end--
			} else if diff == 0 {
				return target
			} else {
				start++
			}
		}
	}

	return target + globalDiff
}

// 17.Letter Combinations of a Phone Number
//第一个想法是直接利用数字来存储一个map
//利用动态规划，也就是递归来实现，注意边界条件和特殊字符
//想办法利用循环来制裁，[]string的字符随着digits字符往后移动而增加
//如果用循环，会有三重循环。1.先找到字符串的数字字符，2.在找到返回字符串里暂存的字符，3.最后找到数字字符对应的char，然后相加
//func LetterCombinations(digits string) []string {
//	if digits == "" {
//		return []string{}
//	}
//	var digitMap = map[string][]string{
//		"2": {"a", "b", "c"},
//		"3": {"d", "e", "f"},
//		"4": {"g", "h", "i"},
//		"5": {"j", "k", "l"},
//		"6": {"m", "n", "o"},
//		"7": {"p", "q", "r", "s"},
//		"8": {"t", "u", "v"},
//		"9": {"w", "x", "y", "z"},
//	}
//
//	return letterRecursive(digitMap, digits, []string{}, 0)
//}
//
//func letterRecursive(digitMap map[string][]string, digits string, strs []string, index int) []string {
//	var returnStrs []string
//	if len(strs) == 0 && len(digits) == 1 {
//		returnStrs = digitMap[string(digits[index])]
//	} else if len(strs) != 0 && len(strs[0])+1 == len(digits) {
//		for _, str := range strs {
//			for _, char := range digitMap[string(digits[index])] {
//				returnStr := str + char
//				returnStrs = append(returnStrs, returnStr)
//			}
//		}
//	} else {
//		var nextStrs []string
//		for _, char := range digitMap[string(digits[index])] {
//			if len(strs) == 0 {
//				returnStrs = letterRecursive(digitMap, digits, digitMap[string(digits[index])], index+1)
//			} else {
//				for _, str := range strs {
//					nextStr := str + char
//					nextStrs = append(nextStrs, nextStr)
//					returnStrs = letterRecursive(digitMap, digits, nextStrs, index+1)
//				}
//			}
//		}
//	}
//	return returnStrs
//}
func LetterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var digitMap = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	returnStrs := []string{""}
	for _, number := range digits {
		tmpStrs := []string{}
		for _, str := range returnStrs {
			for _, char := range digitMap[string(number)] {
				tmpStrs = append(tmpStrs, str+char)
			}
		}
		returnStrs = tmpStrs
	}
	return returnStrs
}

//18.4Sum
//时间复杂度只能为O(n*n*n)?
//可以用hashmap来存两个数之和，时间复杂度最好为O(n*n)，最坏为O(n*n*n),为了适应[0,0,0,0]这种情况，hash碰撞增大
//func fourSum(nums []int, target int) [][]int {
//	length := len(nums)
//	if length < 4 {
//		return [][]int{}
//	}
//	result := [][]int{}
//	sort.Ints(nums)
//
//	for i, num := range nums {
//		if i != 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		threeNumTarget := target - num
//		for j := i + 1; j < length; j++ {
//			if j != i+1 && nums[j] == nums[j-1] {
//				continue
//			}
//			twoNumTarget := threeNumTarget - nums[j]
//			start := j + 1
//			end := length - 1
//			for start < end {
//				if start != j+1 && nums[start] == nums[start-1] {
//					start++
//					continue
//				}
//				if end != length-1 && nums[end] == nums[end+1] {
//					end--
//					continue
//				}
//				tmpTarget := nums[start] + nums[end]
//				if tmpTarget > twoNumTarget {
//					end--
//				} else if tmpTarget < twoNumTarget {
//					start++
//				} else {
//					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
//					start++
//					end--
//				}
//			}
//		}
//	}
//	return result
//}
func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	cache := map[int][][]int{}
	for i, _ := range nums {
		for j := i + 1; j < length; j++ {
			key := nums[i] + nums[j]
			if values, ok := cache[key]; ok {
				if nums[i] == nums[values[len(values)-1][0]] {
					values[len(values)-1][0] = i
					values[len(values)-1][1] = j
				} else {
					values = append(values, []int{i, j})
					cache[key] = values
				}
			} else {
				cache[key] = [][]int{{i, j}}
			}
		}
	}

	for i, _ := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if values, ok := cache[target-nums[i]-nums[j]]; ok {
				for _, value := range values {
					if j >= value[0] {
						continue
					}
					tmp := make([]int, 0, 4)
					tmp = append(tmp, nums[i], nums[j], nums[value[0]], nums[value[1]])
					result = append(result, tmp)

				}
			}
		}
	}

	return result
}
