package _25_merge_order_list

import "github.com/hq-cml/sward2offer/common"

func Merge(head1, head2 *common.ListNode) *common.ListNode {
    if head1 == nil {
        return head2
    }
    if head2 == nil {
        return head1
    }

    //处理确定第一个节点
    var head *common.ListNode
    if head1.Val <= head2.Val {
        head = head1
        head1 = head.Next
    } else {
        head = head2
        head2 = head2.Next
    }

    //剩余节点处理
    p := head
    for head1 != nil && head2 != nil {
        if head1.Val <= head2.Val {
            p.Next = head1
            head1 = head1.Next
        } else {
            p.Next = head2
            head2 = head2.Next
        }
        p = p.Next
    }

    if head1 != nil {
        p.Next = head1
    } else {
        p.Next = head2
    }

    return head
}

//递归
func MergeRecurs(head1, head2 *common.ListNode) *common.ListNode {
    if head1 == nil {
        return head2
    }
    if head2 == nil {
        return head1
    }

    var head *common.ListNode
    if head1.Val <= head2.Val {
        head = head1
        head1 = head.Next
    } else {
        head = head2
        head2 = head2.Next
    }

    //递归处理余下部分
    head.Next = MergeRecurs(head1, head2)
    return head
}