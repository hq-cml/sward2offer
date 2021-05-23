package common

import (
    "fmt"
    "testing"
)

func TestListNode_Foreach(t *testing.T) {
    type args struct {
        f func(*ListNode)
    }
    tests := []struct {
        name   string
        args   args
    }{
        {
            name:   "case1",
            args:   args{
                f: NodePrint,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var l *ListNode
            l = l.PushNode(2)
            l = l.PushNode(4)
            l = l.PushNode(5)
            l = l.PushNode(8)
            l.Foreach(tt.args.f)
            v, l, b := l.PopNode()
            if b != true {
                fmt.Println("err:", b)
            }
            fmt.Println("Pop:", v)
            l.Foreach(tt.args.f)
        })
    }
}