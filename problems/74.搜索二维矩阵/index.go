package index

// todo 还可以用二分查找确定位置
func searchMatrix(matrix [][]int, target int) bool {
	rowLength := len(matrix)
	if rowLength == 0 {
		return false
	}
	columnLength := len(matrix[0])
	if columnLength == 0 {
		return false
	}

	i := 0
	j := columnLength - 1
	for i <= rowLength-1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
