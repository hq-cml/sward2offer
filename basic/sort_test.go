package basic

import (
    "fmt"
    "testing"
)

func TestPartition(t *testing.T) {
    type args struct {
        arr    []int
        begIdx int
        endIdx int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "case1",
            args: args{
                arr:    []int{1,2,3,4,5,6,7,8},
                begIdx: 0,
                endIdx: 7,
            },
            want: 0,
        },
        {
            name: "case2",
            args: args{
                arr:    []int{8,7,6,5,4,3,2,1},
                begIdx: 0,
                endIdx: 7,
            },
            want: 7,
        },
        {
            name: "case3",
            args: args{
                arr:    []int{5,1,3,4,2,7,8,6},
                begIdx: 0,
                endIdx: 7,
            },
            want: 4,
        },
        {
            name: "case4",
            args: args{
                arr:    []int{5,1,3,4,2,7,6},
                begIdx: 0,
                endIdx: 6,
            },
            want: 4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Partition(tt.args.arr, tt.args.begIdx, tt.args.endIdx); got != tt.want {
                t.Errorf("Partition() = %v, want %v", got, tt.want)
            }
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
        //{
        //    name: "case1",
        //    args: args{
        //        arr: []int{5,1,3,4,2,7,8,6},
        //    },
        //},
        //{
        //    name: "case2",
        //    args: args{
        //        arr: []int{5,1,3,4,2,7,6},
        //    },
        //},
        //{
        //    name: "case3",
        //    args: args{
        //        arr: []int{1,2,3,4,5,6,7,8},
        //    },
        //},
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