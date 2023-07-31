/*
 * 求最长的回文子字符串
 */
package _005_longest_palindrome

// 这题的思路和leetcode 647类似
func Longest(s string) string {
	if len(s) <= 1 {
		return s
	}

	length := len(s)
	dp:=make([][]bool, length)
	for i:=0; i<length; i++ {
		dp[i] = make([]bool, length)
	}
	max := 0
	beg := -1
	end := -1
	for i:=length-1; i>=0; i-- {
		for j:=i; j<length; j++ {
			if s[i] != s[j] {
				continue
			}
			if i==j || i==j-1 || dp[i+1][j-1] {
				dp [i][j] = true
				if j-i+1 > max {
					max = j-i+1
					beg, end = i, j
				}
			}
		}
	}
	return s[beg:end+1]
}