package index

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	quickSort(intervals, 0, length-1)
	result := [][]int{}
	var start, end int
	for i := 0; i <= length-1; i++ {
		start = intervals[i][0]
		end = intervals[i][1]
		for i <= length-2 && end >= intervals[i+1][0] {
			if intervals[i+1][1] > end {
				end = intervals[i+1][1]
			}
			i++
		}
		result = append(result, []int{start, end})
	}

	return result
}

func quickSort(intervals [][]int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(intervals, start, end)
	quickSort(intervals, start, mid-1)
	quickSort(intervals, mid+1, end)
}

func partition(intervals [][]int, start, end int) int {
	flag := intervals[end]
	j := start
	for i := start; i <= end-1; i++ {
		if intervals[i][0] < flag[0] {
			intervals[j], intervals[i] = intervals[i], intervals[j]
			j++
		}
	}
	intervals[j], intervals[end] = intervals[end], intervals[j]

	return j
}
