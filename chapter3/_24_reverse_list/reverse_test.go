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