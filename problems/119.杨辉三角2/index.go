package index

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := []int{}
	for k := 0; k <= rowIndex; k++ {
		tmpArray := make([]int, 0, k)
		if k == 0 {
			tmpArray = append(tmpArray, 1)
		} else {
			for i := 0; i <= k; i++ {
				if i == 0 || i == k {
					tmpArray = append(tmpArray, 1)
				} else {
					tmpArray = append(tmpArray, result[i-1]+result[i])
				}
			}
		}
		result = tmpArray
	}

	return result
}
