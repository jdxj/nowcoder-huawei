package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jdxj/nowcoder-huawei/question"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := buf.ReadLine()
		if len(data) <= 0 {
			break
		}
		num, _ := strconv.Atoi(string(data))

		var params []int
		for i := 0; i < num; i++ {
			data, _, _ := buf.ReadLine()
			param, _ := strconv.Atoi(string(data))
			params = append(params, param)
		}

		result := question.ObviouslyRandomNumbers(params)
		for _, v := range result {
			fmt.Println(v)
		}
	}
}
