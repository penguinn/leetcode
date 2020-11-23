package index

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
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

func merge1(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	node := intervals[0]
	tmp := []int{}
	for i := 0; i <= length-1; i++ {
		tmp = intervals[i]
		if tmp[0] >= node[0] && tmp[0] <= node[1] && tmp[1] > node[1] {
			node[1] = tmp[1]
		}else if  tmp[0] >= node[0] && tmp[0] <= node[1] && tmp[1] <= node[1] {
			continue
		} else {
			result = append(result, node)
			node = tmp
		}
	}

	result = append(result, node)

	return result
}
