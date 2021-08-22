package basic

import (
    "fmt"
    "testing"
)

func TestBubbleSort(t *testing.T) {
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
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{5,1,3,4,2,7,6},
            },
        },
        {
            name: "case3",
            args: args{
                arr: []int{1,2,3,4,5,6,7,8},
            },
        },
        {
            name: "case4",
            args: args{
                arr: []int{8,7,6,5,4,3,2,1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            BubbleSort(tt.args.arr)
            fmt.Println(tt.args.arr)
        })
    }
}


func TestQuickSort(t *testing.T) {
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
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{5,1,3,4,2,7,6},
            },
        },
        {
            name: "case3",
            args: args{
                arr: []int{1,2,3,4,5,6,7,8},
            },
        },
        {
            name: "case4",
            args: args{
                arr: []int{8,7,6,5,4,3,2,1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            QuickSort(tt.args.arr)
            fmt.Println(tt.args.arr)
        })
    }
}

func TestHeapSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := HeapSort(tt.args.arr)
            fmt.Println(got)
        })
    }
}