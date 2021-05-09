package _14_cut_rope

import "testing"

func TestMaxCut(t *testing.T) {
    type args struct {
        length int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "case1",
            args: args{
                length: 8,
            },
            want: 18,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MaxCut(tt.args.length); got != tt.want {
                t.Errorf("MaxCut() = %v, want %v", got, tt.want)
            }
        })
    }
}