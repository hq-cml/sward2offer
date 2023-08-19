/*
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度
 * eg1: "(()" => "()" => 2
 * eg2: ")()())" => "()()" =>4
 */
package _032_maxlen_parenthesis

import (
	"fmt"
	"math"
)

// 思路：动态规划
// 核心套路：大问题的解决，依赖于小问题的结果！
// 引入dp[]，dp[i]表示0~i字符子串(以arr[i]结尾的)，表示的最长有效括号子串长度，则：
//
//	如果arr[i] == '('，dp[i]必然为0，因为有效括号不会以(结尾
//	如果arr[i] == ')'，则需要分情况讨论
//	  1. arr[i-i] == '('，则arr[i-1]和arr[i]成功配对，则dp[i] = dp[i-2]+2，这相当于dp[i]的值，依赖于dp[i-2]
//	  2. arr[i-i] == ')'，这种情况相对复杂，可能会形成嵌套匹配，也可能是压根就不匹配
//	     如果dp[i-1]!=0，则需要跳过这个长度，查看是否能够形成嵌套匹配
func CalcMax(arr []byte) int {
	if len(arr) < 1 {
		return 0
	}
	dp := make([]int, len(arr))
	dp[0] = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == '(' {
			continue
		}
		// arr[i] == ')'
		if arr[i-1] == '(' {
			if i-2 >= 0 { // 判断越界
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2 // i==1
			}
		} else {
			// arr[i-1] == ')'
			skipLen := dp[i-1] // 以arr[i-1]作为结尾的最长括号串
			// 直接向前跳过这一段skipLen，取出开头，看下arr[i]是否匹配和
			begIdx := i - skipLen - 1
			if begIdx < 0 {
				dp[i] = 0
			} else if arr[begIdx] == '(' {
				// 形成了嵌套匹配
				dp[i] = dp[i-1] + 2
				// 继续往前追溯，如果前面也是合法字符串则继续追加
				if begIdx-1 >= 0 {
					dp[i] = dp[i] + dp[begIdx-1]
				}
			} else {
				// 不形成匹配
				dp[i] = 0
			}
		}
	}
	fmt.Println(dp)
	max := math.MinInt64
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
