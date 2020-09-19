package index

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	if length < 2 {
		return []int{}
	}
	p1 := 0
	p2 := length - 1

	for p1 != p2 {
		if numbers[p1]+numbers[p2] == target {
			return []int{p1 + 1, p2 + 1}
		} else if numbers[p1]+numbers[p2] < target {
			p1++
		} else {
			p2--
		}
	}
	return []int{p1 + 1, p2 + 1}
}
