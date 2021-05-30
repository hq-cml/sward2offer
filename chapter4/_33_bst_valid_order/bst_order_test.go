package _33_bst_valid_order

import "testing"

func TestCheckValidOrder(t *testing.T) {
    type args struct {
        order []int
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
           name: "case1",
           args: args{
               order: []int{5,7,6,9,11,10,8},
           },
           want: true,
        },
        {
            name: "case2",
            args: args{
                order: []int{7,4,6,5},
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CheckValidOrder(tt.args.order); got != tt.want {
                t.Errorf("CheckValidOrder() = %v, want %v", got, tt.want)
            }
        })
    }
}