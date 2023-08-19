/*
 * 面试题42：连续子数组的最大和
 * 题目：输入一个整型数组，数组里有正数也有负数。数组中一个或连续的多个整
 * 数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。
 *
 * PS：这道题也是编程之美2.14
 *     leetcode 53
 */
package _42_max_sum_sub_arr

import (
	"math"
)

// 思路1：
// 主要是举一个列子，比如{1,2,-4,3,1,2,-4}，列出一张表格，
// 可以看出，一旦累计和出现了负数，那还不如下一个从头开始更划算。
// 同时，设置两个变量，表示累计值和曾经出现的最大值。
// 难度:4*
func CalcMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := math.MinInt64 //最大值
	accumulate := 0      //累计值
	for _, v := range arr {
		if accumulate <= 0 {
			accumulate = v //accumulate清零，从当前开始，重新算
		} else {
			accumulate += v
		}

		if accumulate > max { //记录出现过的最大值
			max = accumulate
		}
	}
	return max
}

// 思路2：
// 这题其实更通用的是动态规划的思路！！但是真的不容易能想到。
// 通常，用动态规划的思路，去递归考虑问题，但是，最终会选择用循环的思路来实现。
// f(i)表示第i个数字结尾的子数组最大和，则
//
//	f(i)= pData[i]          #i=0 或者 f(i-1)<=0，原因，如果是第一个，或者前面的是负数，那还不如从i开始
//	f(i)= f(i-1) + pData[i] #i!=0 且 f(i-1)>0，和上面正好相对应，前面的累计是整数，那自然将i算上更好
//
// 这道题的动态规划的代码，和上面的代码是一样的，只是f(i)对应变量accumulate，max[f(i)]对应max
// 难度:5*
func CalcDynamic(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := math.MinInt64 //目前的最大值
	fI := 0              //f(i)，累计值
	for i, v := range arr {
		if i == 0 || fI <= 0 {
			fI = v //从当前开始，从头开始算
		} else {
			fI = fI + v
		}

		if fI > max { //记录出现过的最大值
			max = fI //这里也可以用一个数组，将每个fI都存下来，最后统一max，貌似更加直观一些，但是空间复杂度就大了
		}
	}
	return max
}
