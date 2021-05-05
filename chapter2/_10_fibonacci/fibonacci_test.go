package _10_fibonacci

import "testing"

func TestFibonacciNoRecurse(t *testing.T) {
    type args struct {
        n int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "case1",
            args: args{
                n : 3,
            },
            want: FibonacciRecurse(3),
        },
        {
            name: "case2",
            args: args{
                n : 10,
            },
            want: FibonacciRecurse(10),
        },
        //{
        //    name: "case3",
        //    args: args{
        //        n : 50,
        //    },
        //    want: FibonacciRecurse(50),
        //},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FibonacciNoRecurse(tt.args.n); got != tt.want {
                t.Errorf("FibonacciNoRecurse() = %v, want %v", got, tt.want)
            }
        })
    }
}