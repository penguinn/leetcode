package index

// 时间复杂度O(n)， 空间复杂度O(n)
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{k * shorter}
	}
	list := make([]int, 0, k+1)
	for i := k; i >= 0; i-- {
		list = append(list, i*shorter+(k-i)*longer)
	}
	return list
}
