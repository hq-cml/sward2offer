/*
 * 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天
 * 下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
 * 示例 1:
 * 输入: temperatures = [73,74,75,71,69,72,76,73]
 * 输出: [1,1,4,2,1,1,0,0]
 * 示例 2:
 *
 * 输入: temperatures = [30,40,50,60]
 * 输出: [1,1,1,0]
 * 示例 3:
 *
 * 输入: temperatures = [30,60,90]
 * 输出: [1,1,0]
 */
package _739_temperature

// 思路：单调栈（从栈底部开始向上，数字越来越大或者越来越小）
func DailyTemperatures(temperatures []int) []int {
	var stk []int // 单调栈
	ret := make([]int, len(temperatures))

	// 遍历温度
	for idx, t := range temperatures {
		// 若当日温度>栈顶温度，说明当日是栈顶日的升温日
		for len(stk) != 0 && temperatures[stk[len(stk)-1]] < t {
			topIdx := stk[len(stk)-1]
			// 记录栈顶日的结论，当日日期 - 栈顶日期
			ret[topIdx] = idx - topIdx
			// 栈顶日记录完毕之后，pop出栈
			stk = stk[:len(stk)-1] // pop
		}

		stk = append(stk, idx) // push
	}
	return ret
}
