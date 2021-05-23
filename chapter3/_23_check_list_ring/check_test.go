package _23_check_list_ring

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestFindEntry(t *testing.T) {
    tests := []struct {
        name string
        want *common.ListNode
    }{
        {
            name: "case1",
            want: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var l *common.ListNode
            l = l.PushNode(1)
            l = l.PushNode(2)
            l = l.PushNode(3)
            node := &common.ListNode{
                Val:  4,
            }
            l = l.PushNodePtr(node)
            l = l.PushNode(5)
            l = l.PushNode(6)
            node2 := &common.ListNode{
                Val:  7,
                Next: node,
            }
            l = l.PushNodePtr(node2)

            p := FindEntry(l)
            fmt.Println(p.Val)
        })
    }
}