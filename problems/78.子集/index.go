package index

func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}
	if length == 1 {
		return [][]int{{}, nums}
	}
	result := [][]int{{}}
	tmpResult := [][]int{}
	for i := 0; i <= length-1; i++ {
		tmpResult = [][]int{}
		for _, e := range result {
			tmp := make([]int, len(e)+1)
			copy(tmp, e)
			tmp[len(e)] = nums[i]
			tmpResult = append(tmpResult, tmp)
		}
		result = append(result, tmpResult...)
	}

	return result
}
