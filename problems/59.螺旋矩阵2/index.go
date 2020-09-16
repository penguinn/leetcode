package index

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	matrix := [][]int{}
	for i := 0; i <= n-1; i++ {
		matrix = append(matrix, make([]int, n))
	}
	upLevel, leftLevel := 0, 0
	downLevel, rightLevel := n-1, n-1
	num := 1
	for {
		for i := leftLevel; i <= rightLevel; i++ {
			matrix[upLevel][i] = num
			num++
		}
		upLevel++
		if upLevel > downLevel {
			break
		}
		for i := upLevel; i <= downLevel; i++ {
			matrix[i][rightLevel] = num
			num++
		}
		rightLevel--
		if rightLevel < leftLevel {
			break
		}
		for i := rightLevel; i >= leftLevel; i-- {
			matrix[downLevel][i] = num
			num++
		}
		downLevel--
		if downLevel < upLevel {
			break
		}
		for i := downLevel; i >= upLevel; i-- {
			matrix[i][leftLevel] = num
			num++
		}
		leftLevel++
		if leftLevel > rightLevel {
			break
		}
	}

	return matrix
}
