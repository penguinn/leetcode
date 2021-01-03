package index

func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)
	if length == 0 || k == 0 || k > length {
		return []int{}
	}
	for i := k/2 - 1; i >= 0; i-- {
		heap(arr, i, k-1)
	}
	for i := k; i <= length-1; i++ {
		if arr[i] < arr[0] {
			arr[0], arr[i] = arr[i], arr[0]
			heap(arr, 0, k-1)
		}
	}
	return arr[:k]
}

// 大顶堆
func heap(arr []int, start, end int) {
	left := 2*start + 1
	right := 2*start + 2
	if left > end {
		return
	}
	tmp := left
	if right <= end && arr[right] > arr[left] {
		tmp = right
	}
	if arr[tmp] > arr[start] {
		arr[tmp], arr[start] = arr[start], arr[tmp]
		heap(arr, tmp, end)
	}
	return
}
