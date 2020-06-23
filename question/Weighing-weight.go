package question

func WeighingWeight(weight, count []int) int {
	var total int
	for i, kind := range weight {
		total += kind * count[i]
	}
	dp := make([]bool, total+1)

	dp[0] = true
	for i := 0; i < len(weight); i++ {
		for j := 0; j < count[i]; j++ {
			for k := total - weight[i]; k >= 0; k-- {
				if dp[k] {
					dp[k+weight[i]] = true
				}
			}
		}
	}

	var kinds int
	for _, v := range dp {
		if v {
			kinds++
		}
	}
	return kinds
}
