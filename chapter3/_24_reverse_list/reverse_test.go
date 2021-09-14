package _24_reverse_list

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestReverseList(t *testing.T) {
    type args struct {
        head *common.ListNode
    }
    var l *common.ListNode
    l = l.PushNode(2)
    l = l.PushNode(4)
    l = l.PushNode(5)
    l = l.PushNode(8)

    tests := []struct {
        name string
        args args
        want *common.ListNode
    }{
        {
            name: "case1",
            args: args{
                head: l,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ReverseList(tt.args.head)
            got.Foreach(common.NodePrint)
        })
    }
}

//测试环，会导致死循环
//func TestReverseList2(t *testing.T) {
//    type args struct {
//        head *common.ListNode
//    }
//    p1 := &common.ListNode{
//        Val:  1,
//        Next: nil,
//    }
//    p2 := &common.ListNode{
//        Val:  2,
//        Next: nil,
//    }
//    p3 := &common.ListNode{
//        Val:  3,
//        Next: nil,
//    }
//    p4 := &common.ListNode{
//        Val:  4,
//        Next: nil,
//    }
//    p5 := &common.ListNode{
//        Val:  5,
//        Next: nil,
//    }
//    p6 := &common.ListNode{
//        Val:  6,
//        Next: nil,
//    }
//    p1.Next = p2
//    p2.Next = p3
//    p3.Next = p4
//    p4.Next = p5
//    p5.Next = p6
//    p6.Next = p3
//
//
//    tests := []struct {
//        name string
//        args args
//        want *common.ListNode
//    }{
//        {
//            name: "case1",
//            args: args{
//                head: p1,
//            },
//        },
//    }
//    for _, tt := range tests {
//        t.Run(tt.name, func(t *testing.T) {
//            _ = ReverseList(tt.args.head)
//            //got.Foreach(common.NodePrint)
//        })
//    }
//}