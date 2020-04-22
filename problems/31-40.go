package problems

import (
	"sort"
)

//40.Combination Sum II
//这道题用的深度优先遍历，时间复杂度为O(2的n次方)
func combinationSum2(candidates []int, target int) [][]int {
	if target < 0 {
		return [][]int{}
	}
	for _, i := range candidates {
		if i < 0 {
			return [][]int{}
		}
	}

	if len(candidates) == 1 {
		if candidates[0] == target {
			return [][]int{candidates}
		} else {
			return [][]int{}
		}
	}

	sort.Ints(candidates)

	result := [][]int{}
	combinationSumReverse2(target, candidates, 0, []int{}, &result)
	return result
}

func combinationSumReverse2(target int, res []int, start int, tmpResult []int, result *[][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		(*result) = append((*result), tmpResult)
	} else {
		if len(res) == 0 {
			return
		} else {
			for i, ch := range res[start:] {
				index := i + start
				if index > start && ch == res[index-1] {
					continue
				}
				inResult := make([]int, len(tmpResult)+1)
				copy(inResult, tmpResult)
				inResult[len(tmpResult)] = ch
				combinationSumReverse2(target-ch, res, index+1, inResult, result)
			}
		}
	}
}
