/*
// 面试题51：数组中的逆序对
// 题目：在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组
// 成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

神仙题，可以用朴素的思想，但是神仙算法很难搞出来。
作者的方案是，先拆分成最细粒度子数组，然后利用归并排序逐个算逆序数。。。

HQ：2021-6-13 也不能算是神仙算法，耐心看还是可以看懂的，
本质上就是将逆序数的统计，放在了归并排序中去实现，牛逼，关键是作者的思路是真的清晰，我自己想，肯定是想不出来的
 */
package _51_inverse_pairs_number

import "fmt"

//归并排序
//同时计算逆序数
//难度：5*
func InverseParis(data []int) int {
    if len(data) == 0 {
        return 0
    }
    backup := make([]int, len(data))
    copy(backup, data) //深拷贝

    cnt := inversePairs(data, backup, 0, len(data)-1)
    fmt.Println("X-------", data)
    fmt.Println("Y-------", backup)
    return cnt
}

//递归
//比较烧脑的地方是backup和data的参数位置，是来回调换的
func inversePairs(data, backup []int, start, end int) int{
    //递归的结束条件
    if start == end {
        backup[start] = data[start]
        return 0
    }

    length := (end - start) / 2
    //递归左右
    //比较烧脑的：backup和data的参数位置，是调换了的
    left := inversePairs(backup, data, start, start+length)
    right := inversePairs(backup, data, start+length+1, end)

    //到了此处，递归已经结束，此时data的前半部分和后半部分，已经排序完毕
    //归并排序是针对backup进行处理的，这也是为什么需要调换的原因
    i := start + length
    j := end

    indexCopy := end
    cnt := 0
    //归并处理的同时，计算逆序数
    for i >= start && j >= start+length+1 {
        if data[i] > data[j] {
            backup[indexCopy] = data[i]
            cnt += (j-start-length)
            indexCopy --
            i --
        } else {
            backup[indexCopy] = data[j]
            indexCopy --
            j --
        }
    }

    //剩余的处理，逆序数不会再有变化
    for i >= start {
        backup[indexCopy] = data[i]
        indexCopy --
        i--
    }
    for j >= start+length+1 {
        backup[indexCopy] = data[j]
        indexCopy --
        j --
    }

    fmt.Println("A-------", data[start: end+1])
    fmt.Println("B-------", backup[start: end+1])
    fmt.Println()
    return cnt + left + right
}