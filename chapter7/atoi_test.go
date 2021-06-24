package chapter7

import (
    "fmt"
    "strconv"
    "strings"
    "testing"
)
func TestOther(t *testing.T) {
    s := "  asd  bce  "
    fmt.Println(strings.Trim(s, " "))

    fmt.Println(strconv.Atoi("123"))
    fmt.Println(strconv.Atoi("123.4"))
    fmt.Println(strconv.Atoi("123.5"))
}

func TestAtoi(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name    string
        args    args
        want    int
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                str: "108",
            },
            want:    108,
            wantErr: false,
        },
        {
            name:    "case2",
            args:    args{
                str: "-",
            },
            want:    0,
            wantErr: true,
        },
        {
            name:    "case3",
            args:    args{
                str: " 10 8 ",
            },
            want:    0,
            wantErr: true,
        },
        {
            name:    "case4",
            args:    args{
                str: " 10x8 ",
            },
            want:    0,
            wantErr: true,
        },
        {
            name:    "case5",
            args:    args{
                str: " -10 8 ",
            },
            want:    0,
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Atoi(tt.args.str)
            if (err != nil) != tt.wantErr {
                t.Errorf("Atoi() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Atoi() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestAtoiApproximate(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name    string
        args    args
        want    int
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                str: "108.1",
            },
            want:    108,
            wantErr: false,
        },
        {
            name:    "case2",
            args:    args{
                str: "108.5",
            },
            want:    109,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := AtoiApproximate(tt.args.str)
            if (err != nil) != tt.wantErr {
                t.Errorf("AtoiApproximate() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("AtoiApproximate() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_checkOverflow(t *testing.T) {
    type args struct {
        n     int
        minus bool
    }

    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "case1",
            args: args{
                n:     0x7FFFFFFFFFFFFFFF,
                minus: false,
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := checkOverflow(tt.args.n, tt.args.minus); got != tt.want {
                t.Errorf("checkOverflow() = %v, want %v", got, tt.want)
            }
        })
    }
}