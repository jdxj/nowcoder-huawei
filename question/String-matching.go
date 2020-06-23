package question

import "strings"

func StringMatching(short, long string) bool {
	for _, v := range short {
		if strings.Index(long, string(v)) < 0 {
			return false
		}
	}
	return true
}
