package _59_list_max_value

import (
    "fmt"
    "testing"
)

func TestFindWindowMax(t *testing.T) {
    type args struct {
        arr       []int
        windowLen int
    }
    tests := []struct {
        name  string
        args  args
        want  []int
        want1 bool
    }{
        {
            name:  "case1",
            args:  args{
                arr:       []int{2,3,4,2,6,2,5,1},
                windowLen: 3,
            },
            want1: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1 := FindWindowMax(tt.args.arr, tt.args.windowLen)
            if got1 != tt.want1 {
                t.Errorf("FindWindowMax() got1 = %v, want %v", got1, tt.want1)
            }
            fmt.Println(got)
        })
    }
}

func TestFindWindowMaxSlide(t *testing.T) {
    type args struct {
        arr       []int
        windowLen int
    }
    tests := []struct {
        name  string
        args  args
        want  []int
        want1 bool
    }{
        {
            name:  "case1",
            args:  args{
                arr:       []int{2,3,4,2,6,2,5,1},
                windowLen: 3,
            },
            want1: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, _ := FindWindowMaxSlide(tt.args.arr, tt.args.windowLen)
            fmt.Println(got)
        })
    }
}