/*
 // 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。

第一反应想到了链表翻转，但是这会改变链表结构，通常来说，面试官不会同意。然后我想到了利用栈，但是要自己实现一个栈？貌似不太现实。
这时候，应该想到栈的本质，其实，递归就是利用了栈，不是吗。所以想到了一种递归思路！！！
所以，如果不方便用栈的时候，考虑是否可以用递归来代替！！
 */
package _06_revert_print_list

import (
    "fmt"
)

type Node struct {
    I int
    next *Node
}

func (l *Node)insert(i int) *Node {
    if l == nil {
        l = &Node{
            I:    i,
            next: nil,
        }
    } else {
        l = &Node{
            I:    i,
            next: l,
        }
    }
    return l
}

func (l *Node)print() {
    p := l
    for p != nil {
        fmt.Print(p.I, ", ")
        p = p.next
    }
    fmt.Println()
}

func RevertPrintList (l *Node) {
    if l == nil {
        return
    }

    RevertPrintList(l.next)
    fmt.Print(l.I, ", ")
}