package _49_ugly_number

import "testing"

func TestGetUglyNumber(t *testing.T) {
    type args struct {
        idx int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "case1",
            args: args{
                idx: 11,
            },
            want: 15,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := GetUglyNumber(tt.args.idx); got != tt.want {
                t.Errorf("GetUglyNumber() = %v, want %v", got, tt.want)
            }
        })
    }
}