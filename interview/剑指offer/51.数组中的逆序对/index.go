package index

var result int
var tmp []int

func reversePairs(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	result = 0
	tmp = make([]int, length)
	mergeSort(nums, 0, length-1)
	return result
}

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	if start >= end {
		return
	}
	i := start
	j := mid + 1
	k := start
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
			k++
		} else {
			tmp[k] = nums[j]
			result += mid - i + 1
			j++
			k++
		}
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= end {
		tmp[k] = nums[j]
		j++
		k++
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i]
	}

	return
}
