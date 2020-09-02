package index

func spiralOrder(matrix [][]int) []int {
	len1 := len(matrix)
	if len1 == 0 {
		return []int{}
	}
	if len1 == 1 {
		return matrix[0]
	}
	len2 := len(matrix[0])
	result := make([]int, 0, len1*len2)
	var block1, block2, block3, block4 int
	block3 = len2 - 1
	block4 = len1 - 1
	for block1 <= block3 || block2 <= block4 {
		for i := block1; i <= block3; i++ {
			result = append(result, matrix[block2][i])
		}
		block2++
		if block2 > block4 {
			break
		}
		for i := block2; i <= block4; i++ {
			result = append(result, matrix[i][block3])
		}
		block3--
		if block1 > block3 {
			break
		}
		for i := block3; i >= block1; i-- {
			result = append(result, matrix[block4][i])
		}
		block4--
		if block2 > block4 {
			break
		}
		for i := block4; i >= block2; i-- {
			result = append(result, matrix[i][block1])
		}
		block1++
		if block1 > block3 {
			break
		}
	}

	return result
}
