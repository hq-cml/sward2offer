package _57_sum_is_s

import (
    "fmt"
    "testing"
)

func TestFindTwoNum(t *testing.T) {
    type args struct {
        arr []int
        sum int
    }
    tests := []struct {
        name  string
        args  args
        want  int
        want1 int
        want2 bool
    }{
        {
            name:  "case1",
            args:  args{
                arr: []int{1,2,4,7,11,15},
                sum: 15,
            },
            want:  4,
            want1: 11,
            want2: true,
        },
        {
            name:  "case2",
            args:  args{
                arr: []int{1,2,4,7,11,15,16},
                sum: 20,
            },
            want:  4,
            want1: 16,
            want2: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1, got2 := FindTwoNum(tt.args.arr, tt.args.sum)
            if got != tt.want {
                t.Errorf("FindTwoNum() got = %v, want %v", got, tt.want)
            }
            if got1 != tt.want1 {
                t.Errorf("FindTwoNum() got1 = %v, want %v", got1, tt.want1)
            }
            if got2 != tt.want2 {
                t.Errorf("FindTwoNum() got2 = %v, want %v", got2, tt.want2)
            }
        })
    }
}

func TestFindSequence(t *testing.T) {
    type args struct {
        sum int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                sum: 9,
            },
        },
        {
            name: "case2",
            args: args{
                sum: 15,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FindSequence(tt.args.sum)
            fmt.Println(got)
        })
    }
}