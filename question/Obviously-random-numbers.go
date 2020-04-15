package question

import (
	"sort"
)

// 明明的随机数
func ObviouslyRandomNumbers(params []int) []int {
	filter := make(map[int]struct{})
	for _, v := range params {
		filter[v] = struct{}{}
	}

	result := make([]int, 0, len(filter))
	for k := range filter {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

//func main() {
//	buf := bufio.NewReader(os.Stdin)
//	for {
//		data, _, _ := buf.ReadLine()
//		if len(data) <= 0 {
//			break
//		}
//		num, _ := strconv.Atoi(string(data))
//
//		var params []int
//		for i := 0; i < num; i++ {
//			data, _, _ := buf.ReadLine()
//			param, _ := strconv.Atoi(string(data))
//			params = append(params, param)
//		}
//
//		result := question.ObviouslyRandomNumbers(params)
//		for _, v := range result {
//			fmt.Println(v)
//		}
//	}
//}
