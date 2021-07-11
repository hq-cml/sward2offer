/*
// 面试题50（一）：字符串中第一个只出现一次的字符
// 题目：在字符串中找出第一个只出现一次的字符。如输入"abaccdeff"，则输出
// 'b'。
// 面试题50（二）：字符流中第一个只出现一次的字符
// 题目：请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从
// 字符流中只读出前两个字符"go"时，第一个只出现一次的字符是'g'。当从该字
// 符流中读出前六个字符"google"时，第一个只出现一次的字符是'l'。



我的思路是排序，作者给的思路是空间换时间，引入一个hash表。两次循环解决问题。
并且，这个hash表还是一个小型的优化变种，char[256]

 */

package _50_first_once_char

//第一个只出现一次的字符
//最核心的思路：空间换时间！
func FindFirst(str string) byte {
    m := make([]int, 256)
    for i, _ := range m {
        m[i] = -1
    }

    for pos, c := range str {
        if m[c] == -1 { //第一次出现，存储其位置
            m[c] = pos
        } else if m[c] >= 0 {
            m[c] = -2   //已经出现过一次了，则c不再有效
        } else {
            //已经出现多次，则什么也不敢
        }
    }

    //扫描出只出现1次的，且最先出现的
    minIdx := -1
    var firstChar byte
    for c, idx := range m {
        if idx < 0 {   //idx==-1或者-2，说明未出现或者出现过多次
            continue
        }
        if minIdx == -1 || idx < minIdx {
            minIdx = idx
            firstChar = byte(c)
        }
    }
    return firstChar
}

//书上还有题目二，计算流，思路和第一题是一致的，只是换成面向对象写法