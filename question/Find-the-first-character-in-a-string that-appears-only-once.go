package question

import "fmt"

func FindTheFirstCharacterInAStringThatAppearsOnlyOnce(param string) string {
	record := make(map[int32]int)
	for _, c := range param {
		record[c]++
	}

	for _, c := range param {
		if record[c] == 1 {
			return string(c)
		}
	}
	return fmt.Sprintf("%d", -1)
}
