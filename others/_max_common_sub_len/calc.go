/*
 * 最大的公共子串的长度
 * 求两个串的所有子串中能够匹配上的最大长度是多少。
 * https://blog.csdn.net/blue_mxy/article/details/108087750
 */
package _max_common_sub_len

// 思路：动态规划
// 矩阵思维，利用dp二维数组，假设str1=abca, str2=bca，则dp数组如下
//    a  b  c  a
// b  0  1  0  0
// c  0  0  2  0
// a  1  0  0  1
// 遍历填充这个数组，相当于逐个探测了子串长度
// dp[i][j] = 0, str1[i]!=str2[j]
// dp[i][j] = dp[i-1][j-1] + 1, str1[i]==str2[j]
func MaxCommonSubLen(s1, s2 []byte) int {
	len1 := len(s1)
	len2 := len(s2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	var dp [][]int
	// len1行，len2列
	for i:=0; i<len1; i++ {
		dp = append(dp, make([]int, len2))
	}

	// 循环填充，记录最大值
	max := 0
	for i:=0; i<len1; i++ {
		for j:=0; j<len2; j++ {
			if s1[i] == s2[j] {
				if i== 0 || j == 0 {
					// 如果处在边界，则直接赋值1
					dp[i][j] = 1
				} else {
					// 非边界，则利用之前的预存值
					dp[i][j] = dp[i-1][j-1]+1
				}

				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}