/*
 * 面试题67：把字符串转换成整数
 * 题目：请你写一个函数StrToInt，实现把字符串转换成整数这个功能。当然，不
 * 能使用atoi或者其他类似的库函数。
 */
package chapter7

import (
    "errors"
    "math"
    "strings"
)

//思路：主要是各种边界和异常考虑的全面性
//1. 输入非法字符
//2. 正负号，包括仅仅只有正负号
//3. 越界问题
//4. 不支持浮点小数
func Atoi(str string) (int, error) {
    //过滤空格字符
    str = strings.Trim(str, " ")
    if str == "" {
        return 0, errors.New("Invalid input")
    }

    //正负号
    sign := 1
    if str[0] == '+' {
        str = str[1:]
    } else if string(str[0]) == "-" {
        str = str[1:]
        sign = -1
    }

    //仅有正负号
    if len(str) == 0 {
        return 0, errors.New("Invalid input")
    }

    ret := 0
    for i:=0; i<len(str); i++ {
        if str[i] >= '0' && str[i] <= '9' {
            ret = ret * 10 + int(str[i] - '0')
            if checkOverflow(ret, sign==-1) {
                return 0, errors.New("Over flow")
            }
        } else {
            return 0, errors.New("Invalid input")
        }
    }
    return sign * ret, nil
}

//校验是否溢出
//最大的有符号数（32位）：0x7FFFFFFF == math.MaxInt32
//最小的有符号数（32位）：0x80000000 == math.MinInt32
func checkOverflow(n int, minus bool) bool {
    if !minus {
        return n > math.MaxInt64
    } else {
        return n < math.MinInt64
    }
}

//带四舍五入的Atoi
//支持对于小数的四舍五入
func AtoiApproximate(str string) (int, error) {
    //过滤空格字符
    str = strings.Trim(str, " ")
    if str == "" {
        return 0, errors.New("Invalid input")
    }

    //正负号
    sign := 1
    if str[0] == '+' {
        str = str[1:]
    } else if string(str[0]) == "-" {
        str = str[1:]
        sign = -1
    }

    //仅有正负号
    if len(str) == 0 {
        return 0, errors.New("Invalid input")
    }

    ret := 0
    addFlag := false //小数进位标记
    for i:=0; i<len(str); i++ {
        if str[i] >= '0' && str[i] <= '9' {
            ret = ret * 10 + int(str[i] - '0')
        } else if string(str[i]) == "." {
            if i+1 <= len(str) - 1 && str[i+1] >= '0' && str[i+1] <= '9'{
                if str[i+1] <= '4' {
                    addFlag = false
                } else {
                    addFlag = true
                }
                break
            }else {
                return 0, errors.New("Invalid input")
            }
        } else {
            return 0, errors.New("Invalid input")
        }
    }

    if checkOverflow(ret, sign==-1) {
        return 0, errors.New("Over flow")
    }

    if addFlag {
        return sign * (ret + 1), nil
    }

    return sign * ret, nil
}
