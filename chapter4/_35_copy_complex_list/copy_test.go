package _35_copy_complex_list

import (
    "fmt"
    "testing"
)

func TestCopy(t *testing.T) {
    type args struct {
        src *CompNode
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                src: newList(),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Copy(tt.args.src)
            printList(tt.args.src)
            fmt.Println("------------------")
            printList(got)
        })
    }
}

func newList () *CompNode {
    node1 := &CompNode{
        Val:   1,
    }
    node2 := &CompNode{
        Val:   2,
    }
    node3 := &CompNode{
        Val:   3,
    }
    node4 := &CompNode{
        Val:   4,
    }
    node5 := &CompNode{
        Val:   5,
    }
    node1.Next  = node2
    node2.Next  = node3
    node3.Next  = node4
    node4.Next  = node5

    node1.Radom = node3
    node2.Radom = node5
    node4.Radom = node2

    return node1
}

func printList(p *CompNode) {
    for p != nil {
        v := 0
        if p.Radom != nil {
            v = p.Radom.Val
        }
        fmt.Println(p.Val, "=>", v)
        p = p.Next
    }
}