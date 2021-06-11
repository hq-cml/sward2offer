package _44_digit_in_sequence

import "testing"

func Test_getWidthOfN(t *testing.T) {
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
                n: 3,
            },
            want: 2700,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := getWidthOfN(tt.args.n); got != tt.want {
                t.Errorf("getWidthOfN() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGetDigit(t *testing.T) {
    type args struct {
        idx int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
         name: "case0",
         args: args{
             idx: 10,
         },
         want: 1,
        },
        {
          name: "case1",
          args: args{
              idx: 190,
          },
          want: 1,
        },
        {
         name: "case2",
         args: args{
             idx: 11,
         },
         want: 0,
        },
        {
         name: "case3",
         args: args{
             idx: 5,
         },
         want: 5,
        },
        {
           name: "case4",
           args: args{
               idx: 13,
           },
           want: 1,
        },
        {
           name: "case5",
           args: args{
               idx: 19,
           },
           want: 4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := GetDigit(tt.args.idx); got != tt.want {
                t.Errorf("GetDigit() = %v, want %v", got, tt.want)
            }
        })
    }
}