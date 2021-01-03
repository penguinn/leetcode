package index

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rowLength := len(matrix)
	if rowLength == 0 {
		return false
	}
	columnLength := len(matrix[0])
	if columnLength == 0 {
		return false
	}
	i := 0
	j := columnLength-1
	for i != rowLength && j != -1 {
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
