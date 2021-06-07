package _43_numberof1_between_1_n

import "testing"

func TestCalc(t *testing.T) {
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
                n: 12,
            },
            want: 5,
        },
        {
            name: "case2",
            args: args{
                n: 188,
            },
            want: Calc2(188),
        },
        {
            name: "case3",
            args: args{
                n: 21345,
            },
            want: Calc2(21345),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Calc(tt.args.n); got != tt.want {
                t.Errorf("Calc() = %v, want %v", got, tt.want)
            }
        })
    }
}