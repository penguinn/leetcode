package index

func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length <= 1 {
		return true
	}

	guard := postorder[length-1]
	count := 0
	index := 0
	for i := 0; i <= length-2; i++ {
		if count == 0 && postorder[i] > guard {
			count = 1
			index = i
		} else if count == 1 && postorder[i] < guard {
			count = 2
		}
	}

	if (count == 0 || count == 1) && verifyPostorder(postorder[:index]) && verifyPostorder(postorder[index:length-1]) {
		return true
	} else {
		return false
	}
}
