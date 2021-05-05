package _09_two_stack_list

import (
    "fmt"
    "testing"
)

func TestNewList(t *testing.T) {
    tests := []struct {
        name string
        want *List
    }{
        {
            name: "case1",
            want: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l := NewList()
            l.AppendTail(1)
            l.AppendTail(2)
            l.AppendTail(3)

            fmt.Println(l.DeleteHead())
            fmt.Println(l.DeleteHead())
            fmt.Println(l.DeleteHead())
            fmt.Println(l.DeleteHead())

            l.AppendTail(1)
            l.AppendTail(2)
            l.AppendTail(3)
            l.DeleteHead()
            l.AppendTail(4)
            l.AppendTail(5)
            fmt.Println(l.DeleteHead())
            fmt.Println(l.DeleteHead())
            fmt.Println(l.DeleteHead())
            fmt.Println(l.DeleteHead())
        })
    }
}