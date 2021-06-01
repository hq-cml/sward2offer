package _38_permutation

import "testing"

func TestPermutation(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                str: "abc",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            Permutation(tt.args.str)
        })
    }
}