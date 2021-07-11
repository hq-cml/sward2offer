/*
// 面试题42：连续子数组的最大和
// 题目：输入一个整型数组，数组里有正数也有负数。数组中一个或连续的多个整
// 数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。


思路1：主要是举一个列子，比如{1,2,-4,3,1,2,-4}，列出一张表格，
可以看出，一旦累计和出现了负数，那还不如从头开始更划算。
设置两个变量表示累计值和曾经出现的最大值。

思路2：这题其实更通用的是动态规划的思路！！但是真的不容易能想到。
通常，用动态规划的思路，去规划考虑问题，但是，最终会选择用循环的思路来实现。
这道题的动态规划的代码，和上面的代码是一样的，只是f(i)对应变量accumulate，max[f(i)]对应maxRet

 */
package _42_max_sum_sub_arr

import (
    "errors"
    "math"
)

//举一个列子，比如{1,2,-4,3,1,2,-4}，列出一张表格
//可以看出，一旦累计和出现了负数，那还不如从当前开始，从头开始更划算
func Calc(arr []int) (int, error) {
    if len(arr) == 0 {
        return 0, errors.New("Invalid input")
    }

    max := math.MinInt64     //最大值
    accumulate := 0          //累计值
    for _, v := range arr {
        if accumulate <= 0 {
            accumulate = v   //从当前开始，从头开始算
        } else {
            accumulate += v
        }

        if accumulate > max {
            max = accumulate
        }
    }
    return max, nil
}