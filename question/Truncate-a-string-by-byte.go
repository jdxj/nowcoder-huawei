package question

func TruncateAStringByByte(param string, count int) string {
	paramBytes := []byte(param)
	paramBytes = paramBytes[:count]
	return string(paramBytes)
}
