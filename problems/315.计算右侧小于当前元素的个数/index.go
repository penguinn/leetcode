package index

var result = []int{}
var switchTable = []int{}
var tmpSwitchTable = []int{}
var tmpTable = []int{}

func countSmaller(nums []int) []int {
	length := len(nums)
	if length < 1 {
		return []int{}
	}
	if length == 1 {
		return []int{0}
	}

	result = []int{}
	switchTable = []int{}
	tmpSwitchTable = []int{}
	tmpTable = []int{}

	for i := 0; i <= length-1; i++ {
		result = append(result, 0)
		switchTable = append(switchTable, i)
		tmpTable = append(tmpTable, -1)
		tmpSwitchTable = append(tmpSwitchTable, -1)
	}

	MergeSort(nums, 0, length-1)

	return result
}

func MergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid+1, end)
	Merge(nums, start, mid, end)

	return
}

func Merge(nums []int, start, mid, end int) {
	i := start
	j := mid + 1
	k := start
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmpTable[k] = nums[i]
			tmpSwitchTable[k] = switchTable[i]
			result[switchTable[i]] += k - i
			k++
			i++
		} else {
			tmpTable[k] = nums[j]
			tmpSwitchTable[k] = switchTable[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmpTable[k] = nums[i]
		tmpSwitchTable[k] = switchTable[i]
		result[switchTable[i]] += k - i
		k++
		i++
	}
	for j <= end {
		tmpTable[k] = nums[j]
		tmpSwitchTable[k] = switchTable[j]
		k++
		j++
	}

	for s := start; s <= end; s++ {
		nums[s] = tmpTable[s]
		switchTable[s] = tmpSwitchTable[s]
	}
}
