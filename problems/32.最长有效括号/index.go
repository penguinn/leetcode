package index

func longestValidParentheses(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 0
	}
	max := 0
	stack := []int{-1}
	for i := 0; i <= length-1; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 1 {
				stack[0] = i
				continue
			} else {
				tmp := i - stack[len(stack)-2]
				stack = stack[:len(stack)-1]
				if tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}
