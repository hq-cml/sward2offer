/*
 * 面试题22：链表中倒数第k个结点
 * 题目：输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，
 * 本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，
 * 从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是
 * 值为4的结点。
 */
package _22_find_back_n_node

import (
    "errors"
    "github.com/hq-cml/sward2offer/common"
)

//思路:
//思路1：最朴素的办法是链表反转，然后求出第k个节点。但是这样会变更了链表本身
//思路2：先遍历求链表长度n，然后，然后从头开始遍历n-k个节点，得到的倒数第k个
//思路3：两个指针，p1先想向前遍历k个节点，然后p2从头开始，和p1向后遍历，知道p1到末尾，此时p2指向倒数k节点
//
//主要考察的鲁棒性，也就是健壮性，其实就是考虑的是否全面：
//head不能为空；n不能是小于0的负数；n不能大于整个链表的长度。。。
func Find(head *common.ListNode, k int)(v int, err error) {
    if head == nil {
        return 0, errors.New("Nil head")
    }
    if k <= 0 { //最后一个，是倒数第一个，不是倒数第0个
        return 0, errors.New("Invalid k")
    }

    p1 := head
    p2 := head

    //p1向前走k个节点
    var i int
    for i = 0; i < k-1; i++ {
        if p1 == nil {
            return 0, errors.New("k > len(head)")
        }
        p1 = p1.Next
    }
    if p1 == nil {
        return 0, errors.New("k > len(head)")
    }

    //p1和p2同步前进
    for p1.Next != nil {
        p2 = p2.Next
        p1 = p1.Next
    }

    return p2.Val, nil
}