package question

import "strings"

// 字符串最后一个单词的长度
func TheLengthOfTheLastWordInTheString(params string) int {
	results := strings.Fields(params)
	resLen := len(results)
	if resLen == 0 {
		return 0
	}

	return len(results[resLen-1])
}

// 编译通过, 内存占用高
//func TheLengthOfTheLastWordInTheString(reader io.Reader) int {
//	var result string
//	var tmp []byte
//	buf := make([]byte, 1)
//
//	for {
//		if _, err := reader.Read(buf); err != nil {
//			panic(err)
//		}
//
//		if buf[0] == ' ' || buf[0] == '\n' {
//			if len(tmp) != 0 {
//				result = string(tmp)
//				tmp = tmp[:0]
//			}
//			if buf[0] == '\n' {
//				break
//			}
//		} else {
//			tmp = append(tmp, buf[0])
//		}
//	}
//
//	return len(result)
//}

// 测试通过, 内存占用较高
//func TheLengthOfTheLastWordInTheString(params string) int {
//	var result string
//  var tmp []byte
//
//	for _, c := range params {
//		if c == ' ' || c == '\n' {
//			if len(tmp) != 0 {
//				result = string(tmp)
//				tmp = tmp[:0]
//			}
//		} else {
//			tmp = append(tmp, byte(c))
//		}
//	}
//	return len(result)
//}
