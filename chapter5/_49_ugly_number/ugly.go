/*
 * 面试题49：丑数
 * 题目：我们把只包含因子2、3和5的数称作丑数（Ugly Number）。求按从小到
 * 大的顺序的第1500个丑数。例如6、8都是丑数，但14不是，因为它包含因子7。
 * 习惯上我们把1当做第一个丑数。
 */
package _49_ugly_number

//思路：
//利用空间换时间，存储之前已经计算好的丑数
//动态规划的思路：
//  首先递归的分析问题，然后反过来自底向上解决问题
//  通常会需要辅助空间，存储中间结果
// 写出丑数序列：1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15 ....
// 可以看出来规律：
// 1. 丑数的序列是递增的
// 2. 第i个丑数，必然是由i前面的某个丑数，乘以2 or 3 or 5得来的
// 3. 所以设计3个指针，分别指向2,3,5为因子的目前为止最大的元素
//难度：5*
func GetUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	//空间换时间，保留数据的中间状态，加速
	uglyNumbers := make([]int, n)
	uglyNumbers[0] = 1 //第一个丑数是1
	nextIdx := 1

	p2 := 0 //以2为因子，到目前的最大指向元素
	p3 := 0 //以3为因子，到目前的最大指向元素
	p5 := 0 //以5为因子，到目前的最大指向元素

	for nextIdx < n {
		min := min(uglyNumbers[p2]*2,
			uglyNumbers[p3]*3,
			uglyNumbers[p5]*5)
		uglyNumbers[nextIdx] = min //第nextIdx个丑数

		//p指针递进
		//因为丑数是有序的，所以如果*因子仍然<=min，则可以快速递进（每轮不一定只有一个p递进，可能是多个p都需要递进）
		//这里比较抽象，需要画图来理解
		for uglyNumbers[p2]*2 <= min {
			p2++
		}
		for uglyNumbers[p3]*3 <= min {
			p3++
		}
		for uglyNumbers[p5]*5 <= min {
			p5++
		}
		nextIdx++
	}
	return uglyNumbers[n-1]
}

//三者取最小
func min(a, b, c int) int {
	m := a
	if a > b {
		m = b
	}
	if m < c {
		return m
	}
	return c
}
