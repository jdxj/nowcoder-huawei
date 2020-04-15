package question

import (
	"strings"
)

// 计算字符个数
func CountTheNumberOfCharacters(param1, param2 string) int {
	var count int
	for i := 0; i < len(param1); i++ {
		if strings.EqualFold(string(param1[i]), param2) {
			count++
		}
	}
	// 其他可用库函数
	//strings.Count()
	return count
}
