package temp

import "math"

func PD(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	length := len(s)
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	var ret int
	max := math.MinInt
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] != s[j] {
				continue
			}
			if i == j || i == (j+1) ||
				dp[i+1][j-1] {
				dp[i][j] = true
				ret++
				if j-i > max {
					max = j - i
				}
			}
		}
	}
	return ret
}