package _58_reverse_string

import "testing"

func TestReverseWord(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "case1",
            args: args{
                str: "I am a student.",
            },
            want: "student. a am I",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReverseWord(tt.args.str); got != tt.want {
                t.Errorf("ReverseWord() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSwapLeft(t *testing.T) {
    type args struct {
        str string
        n   int
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "case1",
            args: args{
                str: "abcdefg",
                n:2,
            },
            want: "cdefgab",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SwapLeft(tt.args.str, tt.args.n); got != tt.want {
                t.Errorf("SwapLeft() = %v, want %v", got, tt.want)
            }
        })
    }
}