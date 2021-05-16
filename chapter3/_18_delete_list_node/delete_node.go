package _18_delete_list_node

import "github.com/hq-cml/sward2offer/common"

//O(1)复杂度删除指定节点
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

//删除有序链表的重复节点
func DeleteDuplicate(head *common.ListNode) (*common.ListNode){
    if head == nil || head.Next == nil {
        return head
    }

    p1 := head
    p2 := head.Next
    for {
        for p2 != nil && p1.Val == p2.Val {
            p2 = p2.Next
            p1.Next = p2
        }
        if p2 == nil {
            break
        }
        p1 = p2
        p2 = p2.Next
    }

    return head
}