package in

import (
	"bufio"
	"fmt"
	"os"
)

func Scanln() {
	var param1, param2 string
	var param3 int
	paramCount, err := fmt.Scanln(&param1, &param2, &param3)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	fmt.Printf("param count: %d, param1: %s, param2: %s, param3: %d\n",
		paramCount, param1, param2, param3)
}

func ScanN() {
	var param1, param2 string
	var param3 int
	paramCount, err := fmt.Scan(&param1, &param2, &param3)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	fmt.Printf("param count: %d, param1: %s, param2: %s, param3: %d\n",
		paramCount, param1, param2, param3)
}

func ContinueScan() {
	for {
		var param string
		n, err := fmt.Scan(&param)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			break
		}
		fmt.Printf("param count: %d, param: [%s]\n", n, param)
		if param == "end" {
			fmt.Printf("exit!")
			break
		}
	}
}

func Scanf() {
	var param1, param2 string
	var param3 int
	paramCount, err := fmt.Scanf("%s\n%s %d", &param1, &param2, &param3)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	fmt.Printf("param count: %d, param1: %s, param2: %s, param3: %d\n",
		paramCount, param1, param2, param3)
}

func Sscan() {
	var param1, param2 string
	var param3 int
	paramCount, err := fmt.Sscan("d d 4", &param1, &param2, &param3)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	fmt.Printf("param count: %d, param1: %s, param2: %s, param3: %d\n",
		paramCount, param1, param2, param3)
}

func BufScan() {
	bufScan := bufio.NewScanner(os.Stdin)
	for bufScan.Scan() {
		param := bufScan.Text()
		fmt.Printf("-recive:[%s]", param)
	}
}
