package index

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	for _, num := range popped {
		for len(stack) == 0 || stack[len(stack)-1] != num {
			if len(pushed) == 0 {
				return false
			} else {
				tmp := pushed[0]
				pushed = pushed[1:]
				stack = append(stack, tmp)
			}
		}
		stack = stack[:len(stack)-1]
	}

	return true
}
