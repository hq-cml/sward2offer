package _23_check_list_ring

import "github.com/hq-cml/sward2offer/common"

//返回：是否成环，环的长度
func CheckRing(head *common.ListNode) (bool, int) {
    if head == nil || head.Next == nil {
        return false, 0
    }

    //判断是否存在环
    p1 := head
    p2 := head
    find := false
    for p2 != nil && p2.Next != nil {
        p2 = p2.Next.Next
        p1 = p1.Next
        if p1 == p2 {
            find = true
            break
        }
    }

    if !find {
        return find, 0
    }

    //存在环，求长度
    l := 1
    p1 = p1.Next
    p2 = p2.Next.Next
    for p1 != p2 {
        l ++
        p1 = p1.Next
        p2 = p2.Next.Next
    }
    return find, l
}

//找到入口
func FindEntry(head *common.ListNode) (*common.ListNode){
    exist, length := CheckRing(head)
    if !exist {
        return nil
    }

    p1 := head
    p2 := head
    for i:=0; i<length; i++ {
        p1 = p1.Next
    }

    for p1 != p2 {
        p1 = p1.Next
        p2 = p2.Next
    }
    return p1
}