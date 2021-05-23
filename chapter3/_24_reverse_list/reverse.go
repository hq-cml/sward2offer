package _24_reverse_list

import "github.com/hq-cml/sward2offer/common"

func ReverseList(head *common.ListNode) *common.ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    p1 := head
    p2 := head.Next
    for p2 != nil {
        tmp := p2.Next
        p2.Next = p1
        p1 = p2
        p2 = tmp
    }
    head.Next = nil //第一个节点的Next指针需要置空
    return p1
}