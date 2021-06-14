package _51_inverse_pairs_number

import "testing"

func TestInverseParis(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
          name: "case1",
          args: args{
              arr: []int{7, 5, 6, 4},
          },
          want: 5,
        },
        {
          name: "case2",
          args: args{
              arr: []int{1, 2, 3, 4},
          },
          want: 0,
        },
        {
           name: "case3",
           args: args{
               arr: []int{4, 3, 2, 1},
           },
           want: 6,
        },
        {
          name: "case4",
          args: args{
              arr: []int{7, 5, 6},
          },
          want: 2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := InverseParis(tt.args.arr); got != tt.want {
                t.Errorf("InverseParis() = %v, want %v", got, tt.want)
            }
        })
    }
}