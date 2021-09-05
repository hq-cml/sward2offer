/*
 * 面试题6：从尾到头打印链表
 * 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。
 */
package _06_revert_print_list

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
)

//最朴素的思路：链表翻转，但是这会改变链表结构
//想到链表的遍历需要和打印顺序相反，类似于栈，但是要自己实现一个栈有点麻烦。
//联想到想到栈的本质，其实，递归就是利用了栈，所以想到了一种递归思路！！！
//所以，如果不方便用栈的时候，考虑是否可以用递归来代替！！
func RevertPrintList (l *common.ListNode) {
    if l == nil {
        return
    }

    RevertPrintList(l.Next)
    fmt.Print(l.Val, ", ")
}