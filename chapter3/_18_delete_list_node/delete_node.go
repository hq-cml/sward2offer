package _18_delete_list_node

import "github.com/hq-cml/sward2offer/common"


func DeleteNodeInO1(head *common.ListNode, del *common.ListNode) (*common.ListNode) {
    if head == nil || del == nil {
        return head
    }

    if head == del { //待删除的是头结点
        return head.Next
    } else if del.Next == nil {
        p1 := head
        p2 := head.Next
        for p2 != del {
            p1 = p2
            p2 = p2.Next
        }
        p1.Next = nil
    } else {
        del.Val = del.Next.Val
        del.Next = del.Next.Next
    }
    return head
}