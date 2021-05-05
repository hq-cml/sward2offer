package _09_two_stack_list

import (
    "container/list"
    "fmt"
    "testing"
)

func TestStack_Push(t *testing.T) {

    tests := []struct {
        name   string
    }{
        {
            name:"case1",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := Stack{
                l1: list.List{},
                l2: list.List{},
            }
            s.Push(1)
            s.Push(2)
            s.Push(3)

            fmt.Println(s.Pop())
            s.Push(4)
            fmt.Println(s.Pop())
            fmt.Println(s.Pop())
            fmt.Println(s.Pop())
        })
    }
}