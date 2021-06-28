/*
// 面试题10：斐波那契数列
// 题目：写一个函数，输入n，求斐波那契（Fibonacci）数列的第n项。
 */
package _10_fibonacci

func FibonacciRecurse(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    return FibonacciRecurse(n-1) + FibonacciRecurse(n-2)
}

func FibonacciNoRecurse(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    i2 := 0
    i1 := 1
    for i := 1; i < n; i ++ {
        tmp := i1
        i1 = i2 + i1
        i2 = tmp
    }
    return i1
}