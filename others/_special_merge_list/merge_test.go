package _special_merge_list

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestAddMerge(t *testing.T) {
    type args struct {
        head1 *common.ListNode
        head2 *common.ListNode
    }
    head1 := common.NewList()
    head1 = head1.PushNode(1)
    head1 = head1.PushNode(3)
    head1 = head1.PushNode(5)
    head1 = head1.PushNode(7)
    head1.Foreach(common.NodePrint)

    head2 := common.NewList()
    head2 = head2.PushNode(2)
    head2 = head2.PushNode(4)
    head2 = head2.PushNode(6)
    head2 = head2.PushNode(8)
    head2.Foreach(common.NodePrint)

    tests := []struct {
        name string
        args args
        want *common.ListNode
    }{
        {
            name: "case1",
            args: args{
                head1: head1,
                head2: head2,
            },
            want: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := AddMerge(tt.args.head1, tt.args.head2)
            got.Foreach(common.NodePrint)
        })
    }
}

func TestAddMerge2(t *testing.T) {
    type args struct {
        head1 *common.ListNode
        head2 *common.ListNode
    }
    head1 := common.NewList()
    head1 = head1.PushNode(1)
    head1 = head1.PushNode(3)
    head1 = head1.PushNode(5)
    head1 = head1.PushNode(7)
    head1.Foreach(common.NodePrint)

    head2 := common.NewList()
    head2 = head2.PushNode(2)
    head2 = head2.PushNode(4)
    head2 = head2.PushNode(6)
    head2 = head2.PushNode(8)
    head2 = head2.PushNode(9)
    head2.Foreach(common.NodePrint)

    tests := []struct {
        name string
        args args
        want *common.ListNode
    }{
        {
            name: "case1",
            args: args{
                head1: head1,
                head2: head2,
            },
            want: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := AddMerge(tt.args.head1, tt.args.head2)
            got.Foreach(common.NodePrint)
        })
    }
}