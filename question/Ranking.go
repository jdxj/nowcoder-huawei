package question

import (
	"fmt"
	"sort"
)

func Ranking(score map[int][]string, order int) {
	sorted := make([]int, 0)
	for k := range score {
		sorted = append(sorted, k)
	}
	sort.Ints(sorted)

	if order == 0 {
		for i := len(sorted) - 1; i >= 0; i-- {
			if students, ok := score[sorted[i]]; ok {
				for _, student := range students {
					fmt.Printf("%s %d\n", student, sorted[i])
				}
			}
		}
	}

	if order == 1 {
		for i := 0; i < len(sorted); i++ {
			if students, ok := score[sorted[i]]; ok {
				for _, student := range students {
					fmt.Printf("%s %d\n", student, sorted[i])
				}
			}
		}
	}
}

//func main() {
//	for {
//		var num, order int
//		_, err := fmt.Scan(&num, &order)
//		if err != nil {
//			break
//		}
//
//		score := make(map[int][]string)
//		for i := 0; i < num; i++ {
//			var name string
//			var s int
//			fmt.Scan(&name, &s)
//			score[s] = append(score[s], name)
//		}
//
//		Ranking(score, order)
//
//	}
//
//}
