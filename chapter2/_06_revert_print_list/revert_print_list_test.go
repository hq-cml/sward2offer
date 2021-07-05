package _06_revert_print_list

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestRevertPrintList(t *testing.T) {
    type args struct {
        l *common.ListNode
    }
    var l *common.ListNode
    l = l.PushNode(1)
    l = l.PushNode(2)
    l = l.PushNode(3)
    l = l.PushNode(4)
    l = l.PushNode(5)
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                l: l,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            fmt.Print("Asc: ")
            l.Foreach(common.NodePrint)
            fmt.Print("Desc: ")
            RevertPrintList(tt.args.l)
            fmt.Println()
        })
    }
}