/*
 * 面试题22：链表中倒数第k个结点
 * 题目：输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，
 * 本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，
 * 从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是
 * 值为4的结点。


本体主要考察的鲁棒性，也就是健壮性，其实就是考虑的是否全面：
head不能为空；n不能是小于0的负数；n不能大于整个链表的长度。。。
 */
package _22_find_back_n_node

import (
    "errors"
    "github.com/hq-cml/sward2offer/common"
)

func Find(root *common.ListNode, k int)(v int, err error) {
    if root == nil {
        return 0, errors.New("Nil root")
    }
    if k <= 0 { //最后一个，是倒数第一个，不是倒数第0个
        return 0, errors.New("Invalid k")
    }

    p1 := root
    p2 := root

    var i int
    for i = 0; i < k-1; i++ {
        if p1 == nil {
            return 0, errors.New("k > len(root)")
        }
        p1 = p1.Next
    }

    if p1 == nil {
        return 0, errors.New("k > len(root)")
    }

    for p1.Next != nil {
        p2 = p2.Next
        p1 = p1.Next
    }

    return p2.Val, nil
}