package index

// 使用动态规划的方式
func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	leftMaxList := make([]int, length)
	leftMax := 0
	for i := 0; i <= length-1; i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		}
		leftMaxList[i] = leftMax
	}

	rightMaxList := make([]int, length)
	rightMax := 0
	for i := length - 1; i >= 0; i-- {
		if height[i] > rightMax {
			rightMax = height[i]
		}
		rightMaxList[i] = rightMax
	}

	result := 0
	for i := 1; i <= length-2; i++ {
		var max int
		if leftMaxList[i] > rightMaxList[i] {
			max = rightMaxList[i]
		} else {
			max = leftMaxList[i]
		}
		result += (max - height[i])
	}

	return result
}

// 使用双指针
func trap1(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	result := 0

	left := 0
	right := length - 1
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] < leftMax {
				result += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				result += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}

	return result
}

// 使用栈
// 1. 当前高度小于等于栈顶高度，入栈，指针后移。
// 2. 当前高度大于栈顶高度，出栈，计算出当前墙和栈顶的墙之间水的多少，然后计算当前的高度和新栈的高度的关系，重复第 2 步。直到当前墙的高度不大于栈顶高度或者栈空，然后把当前墙入栈，指针后移。
func trap2(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	result := 0

	stack := []int{}
	current := 0
	var distance, minH int
	for current <= length-1 {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			topHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance = current - stack[len(stack)-1] - 1
			if height[current] > height[stack[len(stack)-1]] {
				minH = height[stack[len(stack)-1]]
			} else {
				minH = height[current]
			}
			result += (minH - topHeight) * distance
		}
		stack = append(stack, current)
		current++
	}

	return result
}


func trap5(height []int) int {
	length := len(height)
	if length <= 0 {
		return 0
	}
	left := 0
	right := length-1
	maxLeft := 0
	maxRight := length-1
	result := 0
	for left < right {
		if height[left] <= height[right] {
			if height[maxLeft] <= height[left] {
				maxLeft = left
			} else {
				result += (height[maxLeft] - height[left])
			}
			left++
		} else {
			if height[maxRight] <= height[right] {
				maxRight = right
			} else {
				result += (height[maxRight] - height[right])
			}
			right--
		}
	}

	return result
}