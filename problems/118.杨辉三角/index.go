package index

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	result := [][]int{}
	for k := 0; k < numRows; k++ {
		if k == 0 {
			result = append(result, []int{1})
		} else {
			array := []int{}
			for i := 0; i <= k; i++ {
				if i == 0 || i == k {
					array = append(array, 1)
				} else {
					array = append(array, result[k-1][i-1]+result[k-1][i])
				}
			}
			result = append(result, array)
		}
	}
	return result
}
