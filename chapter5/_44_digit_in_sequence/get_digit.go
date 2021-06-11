package _44_digit_in_sequence

import "math"

//这个题目相当恶心，因为idx从0开始计数
//所以边界非常烧脑，太烧脑，要保持idx的计数方式统一，否则会乱掉
//难度：5*
func GetDigit(idx int) int {
    n := 1
    //首先判断idx位于几位数
    for {
        if idx - getWidthOfN(n) == 0 {
            return 1 //正好位于n位数的第一个数字的最高位
        }

        if idx - getWidthOfN(n) > 0 {
            idx -= getWidthOfN(n)
            n++
        } else {
            break
        }
    }

    //还剩下的可用的位数，idx仍然从0开始计数，必须要统一
    num := idx / n              //相隔完整的num个数字
    idx = idx % n
    var beg int                 //n位数的开始
    if n == 1 {
        beg = 0
    } else {
        beg = int(math.Pow10(n-1))
    }
    beg = beg + num

    //到这里，就是求beg这个n位数字的第idx位，idx从0开始计数
    cnt := n - idx //倒着求
    ret := 0
    for cnt > 0 {
        ret = beg%10
        beg = beg/10
        cnt --
    }

    return ret
}

//获取N位数的宽度
//比如1位数：10；2位数：180；3位数：2700...
func getWidthOfN(n int) int {
    if n < 1 {
        return 0
    }

    if n == 1 {
        return 10
    }

    return int(math.Pow10(n) - math.Pow10(n-1)) * n
}