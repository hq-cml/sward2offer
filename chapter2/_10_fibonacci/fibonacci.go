/*
 * 面试题10：斐波那契数列
 * 题目：写一个函数，输入n，求斐波那契（Fibonacci）数列的第n项。
 */
package _10_fibonacci

//递归写法
//递归的问题：往往可能存在大量的重复计算！！所以性能不会高！也有可能栈溢出的风险。
func FibonacciRecurse(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    return FibonacciRecurse(n-1) + FibonacciRecurse(n-2)
}

//非递归
//不用递归的时候，自然联想到利用循环或者栈
func FibonacciNoRecurse(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    i1 := 0
    i2 := 1
    for i := 1; i < n; i ++ {
        tmp := i2
        i2 = i1 + i2
        i1 = tmp
    }
    return i2
}