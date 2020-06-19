package index

// 时间复杂度为O(1), 空间复杂度O(n)
func isValid(s string) bool {
	queue := []byte{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			queue = append(queue, byte(c))
		} else if len(queue) <= 0 {
			return false
		} else if c == ')' {
			if queue[len(queue)-1] == '(' {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		} else if c == ']' {
			if queue[len(queue)-1] == '[' {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		} else if c == '}' {
			if queue[len(queue)-1] == '{' {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(queue) == 0 {
		return true
	} else {
		return false
	}
}
