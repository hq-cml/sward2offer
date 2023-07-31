/*
  - 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
  - 有效括号组合需满足：左括号必须以正确的顺序闭合。
    *
    *示例 1：
    输入：n = 3
    输出：["((()))","(()())","(())()","()(())","()()()"]
    示例 2：
    输入：n = 1
    输出：["()"]
*/
package _022_generate_parenthesis

// 简单的回溯解法：
// dfs函数中leftRemain表示当前还有多少左括号可以分配，rightRemain表示当前还有多少右括号可以分配，temp记录当前形成的字符串
// 终止条件：
// 如果leftRemain和rightRemain都为0，即说明所有括号都已加入到字符串，res.add(temp)，返回即可；
// 剪枝（去掉不合法的组合）：
// 如果左括号剩余比右括号多（即形成的字符串中右括号比左括号多）那么不符合规则，直接返回
// 递归：
// 当leftRemain大于0时，还有左括号可以分配，先分配左括号，这是左子树；
// 分配右括号，是右子树，这里由于前面剪枝的判断，所以不会出现rightRemain小于0的情况
func Generate(n int) []string {
	result := []string{}
	// 初始字符串为空，且左Remain和右Remain都是n
	gen(n, n, "", &result)
	return result
}

// 参数s是切片的指针，不是切片
func gen(leftRemain, rightRemain int, str string, s *[]string) {
	// 退出条件：合法组成
	if leftRemain == 0 && rightRemain == 0 {
		*s = append(*s, str)
		return
	}

	// 退出条件：不合法（剪枝）
	if leftRemain > rightRemain {
		//str = str[0 : len(str)-1] //回退一格（这一行可有可无，因为已经被剪枝，不可能存入最终结果）
		return
	}

	// 左子树递归探测
	if leftRemain != 0 {
		gen(leftRemain-1, rightRemain, str+"(", s) // 分配一个左括号
	}

	// 右子树递归探测
	if rightRemain != 0 {
		gen(leftRemain, rightRemain-1, str+")", s) // 分配一个右括号
	}
}

// 2022-12-21，字练习实现
//func gen(leftRemain, rightRemain int, nowStr string, ret *[]string) {
//	// 不符合括号规则，剪枝
//	if leftRemain > rightRemain {
//		return
//	}
//	// 非法情况，剪枝
//	if leftRemain < 0 || rightRemain < 0 {
//		return
//	}
//	// 合法成立
//	if leftRemain == 0 && rightRemain == 0 {
//		*ret = append(*ret, nowStr)
//	}
//
//	// 左右分支递归判断
//	gen(leftRemain - 1, rightRemain, nowStr + "(", ret)
//	gen(leftRemain, rightRemain-1, nowStr + ")", ret)
//}
