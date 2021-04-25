package _06_revert_print_list

import (
    "fmt"
    "testing"
)

func TestRevertPrintList(t *testing.T) {
    type args struct {
        l *Node
    }
    var l *Node
    l = l.insert(1)
    l = l.insert(2)
    l = l.insert(3)
    l = l.insert(4)
    l = l.insert(5)
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
            l.print()
            fmt.Print("Desc: ")
            RevertPrintList(tt.args.l)
        })
    }
}