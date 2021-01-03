package index

func replaceSpace(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	var str string
	for i, value := range s {
		if value == ' ' {
			str += "%20"
		} else {
			str += s[i : i+1]
		}
	}
	
	return str
}
