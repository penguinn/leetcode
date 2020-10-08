package index

func canFinish(numCourses int, prerequisites [][]int) bool {
	numsMap := map[int][]int{}
	for _, prerequisite := range prerequisites {
		_, ok := numsMap[prerequisite[0]]
		if ok {
			numsMap[prerequisite[0]] = append(numsMap[prerequisite[0]], prerequisite[1])
		} else {
			numsMap[prerequisite[0]] = []int{prerequisite[1]}
		}
	}

	flagArray := make([]int, numCourses)
	for i := 0; i <= numCourses-1; i++ {
		ok := dfs(i, flagArray, numsMap)
		if !ok {
			return false
		}
	}

	return true
}

func dfs(index int, flagArray []int, numsMap map[int][]int) bool {
	if flagArray[index] == 1 {
		return false
	} else if flagArray[index] == 2 {
		return true
	} else {
		flagArray[index] = 1
		for _, num := range numsMap[index] {
			ok := dfs(num, flagArray, numsMap)
			if !ok {
				return false
			}
		}
		flagArray[index] = 2
		return true
	}
}
