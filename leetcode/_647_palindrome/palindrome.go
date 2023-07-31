/*
 * 求回文子字符串数量
 * aaa => 6
 * abc => 3
 */
package _647_palindrome

// 动态规划的循环解法，利用一个二维数组，存储回文状态，避免重复计算
// 通过dp二维数组，预存储[i,j]的子串是否是回文子串，布尔类型
//
// 回文判断：
// 一个子字符串str[i,j]是回文，必然是str[i] == str[j]，且满足下面三种情况之一：
// 单个字母：i==j
// 相邻字母：i==j+1
// 不相邻字母：str[i+1, j-1]是回文
func CalcCnt(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := len(s)
	// dp[i][j]表示区间范围[i,j]的子串是否是回文子串，布尔类型
	dp := make([][]bool, length)
	for i:=0; i<length; i++ {
		dp[i] = make([]bool, length)
	}

	var ret int
	// 从最后一个字符开始，逐步向前遍历（注意这是必须的！
	// 不能反过来！想象dp矩阵，如果反过来操作，那么dp[i + 1][j - 1]将是未初始化的值
	for i:= length-1; i>=0; i-- {
		for j:= i; j<length; j++ {
			// s[i] != s[j]，则str[i,j]必然不是回文了
			if s[i] != s[j] {
				continue
			}
			// s[i] == s[j]，到达此处i和j指向的字母已经是相同的了
			if (j - i <= 1) ||	   // 如果i和j指向同一个字母或者i和j相邻，则必然是回文
				dp[i + 1][j - 1] { // 如果i和j不相邻，则查看i+1到j-1是否是回文，如果是则这里也是回文
				ret++;
				dp[i][j] = true;
			}
		}
	}
	return ret
}