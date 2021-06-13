package _50_first_once_char

import "testing"

func TestFindFirst(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name string
        args args
        want byte
    }{
        {
            name: "case1",
            args: args{
                str: "abaccdeff",
            },
            want: byte('b'),
        },
        {
            name: "case2",
            args: args{
                str: "google",
            },
            want: byte('l'),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindFirst(tt.args.str); got != tt.want {
                t.Errorf("FindFirst() = %v, want %v", got, tt.want)
            }
        })
    }
}