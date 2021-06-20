package _59_list_max_value

import (
    "fmt"
    "testing"
)

func TestNewMaxList(t *testing.T) {
    tests := []struct {
        name string
        want *MaxList
    }{
        {
            name: "case1",
            want: NewMaxList(),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := NewMaxList()
            got.Insert(2)
            got.Insert(3)
            got.Insert(4)
            fmt.Println(got.Max())
            fmt.Println(got.Pop())
            fmt.Println(got.Pop())
            fmt.Println(got.Max())
        })
    }
}