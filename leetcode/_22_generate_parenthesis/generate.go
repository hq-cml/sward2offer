/*
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 有效括号组合需满足：左括号必须以正确的顺序闭合。
 *
 *示例 1：
	输入：n = 3
	输出：["((()))","(()())","(())()","()(())","()()()"]
 示例 2：
	输入：n = 1
	输出：["()"]
*/
package _22_generate_parenthesis

//简单的回溯解法：dfs函数中leftRemain表示当前还有多少左括号可以分配，rightRemain表示当前还有多少右括号可以分配，temp记录当前形成的字符串
//终止条件：如果leftRemain和rightRemain都为0，即说明所有括号都已加入到字符串，res.add(temp)，返回即可；
//剪枝：如果当前左括号剩下的比右括号多，即形成的字符串中右括号比左括号多，那么不符合规则，直接返回
//递归：当leftRemain大于0时，还有左括号可以分配，先分配左括号，这是左子树；分配右括号，是右子树，这里由于前面剪枝的判断，所以不会出现rightRemain小于0的情况
func Generate(n int) []string {
	result := []string{}
	gen(n, n, "", &result)
	return result
}

func gen(leftRemain, rightRemain int, str string, s *[]string) {
	if leftRemain == 0 && rightRemain == 0 {
		*s = append(*s, str)
		return
	}

	if leftRemain > rightRemain {
		//不合法
		str = str[0 : len(str)-1] //回退一格
		return
	}

	if leftRemain != 0 {
		gen(leftRemain-1, rightRemain, str+"(", s)
	}

	if rightRemain != 0 {
		gen(leftRemain, rightRemain-1, str+")", s)
	}
}
