package question

func SodaBottle(empty int) int {
	var drink int

	for empty > 2 {
		newEmpty, remain := empty/3, empty%3
		empty = newEmpty + remain

		drink += newEmpty
	}
	if empty == 2 {
		drink++
	}
	return drink
}
