package question

func CountTheNumberOfCapitalLetters(param string) int {
	var count int
	for _, v := range param {
		if 65 <= v && v <= 90 {
			count++
		}
	}
	return count
}
