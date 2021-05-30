package _31_stack_order

import "testing"

func TestCheckValidSequence(t *testing.T) {
    type args struct {
        a []int
        b []int
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "case1",
            args: args{
                a: []int{1,2,3,4,5},
                b: []int{4,5,3,2,1},
            },
            want: true,
        },
        {
            name: "case2",
            args: args{
                a: []int{1,2,3,4,5},
                b: []int{4,3,5,1,2},
            },
            want: false,
        },
        {
            name: "case3",
            args: args{
                a: []int{1,2,3,4,5},
                b: []int{4,3,5,2,1},
            },
            want: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CheckValidSequence(tt.args.a, tt.args.b); got != tt.want {
                t.Errorf("CheckValidSequence() = %v, want %v", got, tt.want)
            }
        })
    }
}