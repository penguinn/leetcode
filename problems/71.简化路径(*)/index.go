package index

import (
	"strings"
)

func simplifyPath(path string) string {
	ret := []string{}
	for _, v := range strings.Split(path, "/") {
		switch v {
		case "":
			break
		case ".":
			break
		case "..":
			if l := len(ret); l > 0 {
				ret = ret[:l-1]
			}
		default:
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}
