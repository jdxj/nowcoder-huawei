package question

import (
	"sort"
	"strings"
)

func TheBeautyOfTheName(name string) int {
	var level int
	count := make([]int, 0)
	tmp := make(map[string]int)
	for _, c := range name {
		v := strings.ToLower(string(c))
		tmp[v]++
	}

	for _, v := range tmp {
		count = append(count, v)
	}
	sort.Ints(count)

	for i := 0; i < len(count); i++ {
		level += count[i] * (26 - len(count) + 1 + i)
	}
	return level
}

//func main() {
//	var num int
//	if _, err := fmt.Scanln(&num); err != nil {
//		panic(err)
//	}
//
//	for i := 0; i < num; i++ {
//		var param string
//		if _, err := fmt.Scanln(&param); err != nil {
//			panic(err)
//		}
//
//		level := question.TheBeautyOfTheName(param)
//		fmt.Println(level)
//	}
//}
