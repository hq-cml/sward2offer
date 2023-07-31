/*
 * 面试题38：字符串的全排列
 * 题目：输入一个字符串，打印出该字符串中字符的所有排列。例如输入字符串abc，
 * 则打印出由字符a、b、c所能排列出来的所有字符串abc、acb、bac、bca、cab和cba。
 */
package _38_permutation

// 思路：
// 总体来说是一个递归思想，深度遍历
// 从字符串的每个字符，分别选出来，占第一位
// 递归得处理剩余的字符串
// 难度：4*
func Permutation(arr []byte) []string {
	var ret []string
	var dfs func(arr []byte, s string)
	dfs = func(arr []byte, s string) {
		// 递归结束条件，字符串长度达到入参长度
		if len(arr) == 0 {
			ret = append(ret, s)
		}
		// 每个字符，分别占第一位
		for i := 0; i < len(arr); i++ {
			c := arr[i]
			// 剩余字符串
			var left []byte
			left = append(left, arr[:i]...)
			left = append(left, arr[i+1:]...)
			// 递归剩余字符串（s + string(c)是到即刻为止字符串的形态）
			dfs(left, s+string(c))
		}
	}
	dfs(arr, "")
	return ret
}
