package _21_swap_arr

import (
    "fmt"
    "testing"
)

func TestSwap1(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{1,2,3,4,5},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{2,2,2,4,2},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            Swap(tt.args.arr, checkOddEvent)
            fmt.Println(tt.args.arr)
        })
    }
}

func TestSwap2(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{1,-2,-3,4,5},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{2,2,2,4,2},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            Swap(tt.args.arr, checkMinus)
            fmt.Println(tt.args.arr)
        })
    }
}