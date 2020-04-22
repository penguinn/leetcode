package problems

import (
	"strings"
	"time"
)

func wordBreak(s string, wordDict []string) bool {
	goalChan := make(chan string, 1000000)
	goalChan <- s
NLoop:
	for {
		select {
		case str := <-goalChan:
			if str == "" {
				return true
			}
			for _, word := range wordDict {
				if strings.HasPrefix(str, word) {
					goalChan <- strings.TrimPrefix(str, word)
				} else {
					continue
				}
			}
		case <-time.After(2 * time.Second):
			break NLoop
		}
	}
	return false
}
