package _45_arrange_array_for_min

import "testing"

func TestArrangeMin(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "case1",
            args: args{
                arr: []int{3, 32, 321},
            },
            want: "321323",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ArrangeMin(tt.args.arr); got != tt.want {
                t.Errorf("ArrangeMin() = %v, want %v", got, tt.want)
            }
        })
    }
}