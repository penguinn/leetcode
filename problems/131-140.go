package problems

//134. Gas Station
//1 使用gas数组来减cost数组，得到的值中最大值作为起点，验证这个点
//func canCompleteCircuit(gas []int, cost []int) int {
//	gasLength := len(gas)
//	costLength := len(cost)
//	if gasLength != costLength || gasLength < 1 {
//		return -1
//	}
//	tmps := make([]int, gasLength)
//	results := make([]int, gasLength)
//	var total int
//	for i, _ := range gas {
//		tmps[i] = gas[i] - cost[i]
//		total += tmps[i]
//	}
//
//	max := math.MinInt64
//	for i, tmp := range tmps {
//		if tmp >= 0 {
//			results[i] = tmp
//			if tmp > max {
//				max = tmp
//			}
//		}
//	}
//	if max < 0 {
//		return -1
//	}
//	result := -1
//	for index, _ := range results {
//		var sum int
//		flag := 1
//		for i := index; i != index+gasLength; i++ {
//			sum = sum + gas[i%gasLength] - cost[i%gasLength]
//			if sum < 0 {
//				flag = 0
//				break
//			}
//		}
//		if flag == 1 {
//			result = index
//			break
//		}
//	}
//	return result
//}

func canCompleteCircuit(gas []int, cost []int) int {
	gasLength := len(gas)
	costLength := len(cost)
	if gasLength != costLength || gasLength < 1 {
		return -1
	}
	var total int
	var tmpTotal int
	var start int
	for i, _ := range gas {
		total += gas[i] - cost[i]
		tmpTotal += gas[i] - cost[i]
		if tmpTotal < 0 {
			tmpTotal = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
