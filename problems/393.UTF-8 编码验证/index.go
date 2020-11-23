package index

func validUtf8(data []int) bool {
	length := len(data)
	if length == 0 {
		return false
	}
	for i:=0; i<=length-1; {
		count := judegeCount(data[i])
		if count == -1 {
			return false
		}
		end := i+count-1
		i++
		for ;i<=end;i++ {
			if i > length-1 || (data[i] & 0xc0 != 0x80) {
				return false
			}
		}
	}

	return true
}

func judegeCount(num int) int {
	if num & 0x80 == 0 {
		return 1
	} else if num & 0xe0 == 0xc0 {
		return 2
	} else if num & 0xf0 == 0xe0 {
		return 3
	} else if num & 0xf8 == 0xf0 {
		return 4
	} else {
		return -1
	}
}
