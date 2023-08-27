/**
 *
 */
package _338_bit_cnt

// 最直观的遍历法
func CountBits(n int) []int {
	if n < 0 {
		return nil
	}
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 0; j < 64; j++ {
			x := 1 << j
			if x > n {
				break
			}
			if i&x != 0 {
				arr[i]++
			}
		}
	}
	return arr
}

// 牛逼方法：动态规划
// dp[i] 表示为值，则
// i为奇数，dp[i] = dp[i-1] + 1
// i为偶数，dp[i] = dp[i/2]
func CountBitsDp(n int) []int {
	if n < 0 {
		return nil
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i/2]
		}
	}
	return dp
}
