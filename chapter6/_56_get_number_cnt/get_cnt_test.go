package _56_get_number_cnt

import "testing"

func TestFindNumerOnlyTwice(t *testing.T) {
    type args struct {
        arr []int
    }
    var tests = []struct {
        name  string
        args  args
        want  int
        want1 int
    }{
        {
            name: "case1",
            args: args{
                arr: []int{2, 4, 3, 6, 3, 2, 5, 5},
            },
            want:  6,
            want1: 4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1 := FindNumerOnlyTwice(tt.args.arr)
            if got != tt.want {
                t.Errorf("FindNumerOnlyTwice() got = %v, want %v", got, tt.want)
            }
            if got1 != tt.want1 {
                t.Errorf("FindNumerOnlyTwice() got1 = %v, want %v", got1, tt.want1)
            }
        })
    }
}

func TestFindNumerOnlyOnce(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "case1",
            args: args{
                arr: []int{1,8, 15, 15, 8, 5, 1,1, 15, 8},
            },
            want: 5,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindNumerOnlyOnce(tt.args.arr); got != tt.want {
                t.Errorf("FindNumerOnlyOnce() = %v, want %v", got, tt.want)
            }
        })
    }
}