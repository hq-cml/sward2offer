/*
 * 面试题21：调整数组顺序使奇数位于偶数前面
 * 题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有
 * 奇数位于数组的前半部分，所有偶数位于数组的后半部分。


这个题目相对简单，其实主要考察的能否写出可扩展的程序，就是把函数作为参数~
 */
package _21_swap_arr

//按规则调换数组
//考虑扩展性：将check函数作为参数
func Swap(arr []int, check func(int)bool) {
    if len(arr) == 0 {
        return
    }

    i := 0
    j := len(arr)-1

    for i<=j {
        if check(arr[i]) {
            i ++
            continue
        }

        if !check(arr[j]) {
            j --
            continue
        }
        //swap
        arr[i], arr[j] = arr[j], arr[i]
    }
}

//奇数返回true，偶数false
func checkOddEvent(n int) bool {
    return n % 2 == 1
}

//正数和0返回true，负数返回false
func checkMinus(n int) bool {
    return n >= 0
}