/*
 * qtt面试题: 另类链表合并，实现类似于加法的操作
 *
 * 例如输入：head1 = [1,3,5,7] head2 = [2,4,6,8]
 * 输出： head = [3, 7, 1, 6, 1]
 * 例如输入：head1 = [1,3,5,7] head2 = [2,4,6,8,9]
 * 输出： head = [3, 7, 1, 6, 0, 1]
 */
package _special_merge_list

import "github.com/hq-cml/sward2offer/common"

func AddMerge(head1 *common.ListNode, head2 *common.ListNode) *common.ListNode{
    if head1 == nil {
        return head2
    }

    if head2 == nil {
        return head1
    }

    result := common.NewList()
    flag := false //进位标记
    p1 := head1
    p2 := head2
    for p1 != nil && p2 != nil {
        tmp := p1.Val + p2.Val
        if flag {
            tmp ++
        }
        if tmp > 9 {
            flag = true
            tmp -= 10
        } else {
            flag = false
        }

        result = result.PushNode(tmp)

        p1 = p1.Next
        p2 = p2.Next
    }

    //p1或者p2，很有可能还存在剩余
    if p1 != nil || p2 != nil {
        var final *common.ListNode
        if p1 != nil {
            final = p1
        } else {
            final = p2
        }

        for final != nil {
            tmp := final.Val
            if flag {
                tmp ++
            }
            if tmp > 9 {
                flag = true
                tmp -= 10
            } else {
                flag = false
            }

            result = result.PushNode(tmp)
            final = final.Next
        }
    }

    //还剩余一个进位
    if flag {
        result = result.PushNode(1)
    }

    return result
}