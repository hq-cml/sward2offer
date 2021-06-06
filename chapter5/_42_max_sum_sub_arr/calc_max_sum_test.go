package _42_max_sum_sub_arr

import "testing"

func TestCalc(t *testing.T) {
    type args struct {
        arr []int
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
                arr: []int{1,2,-4,3,1,2,-4},
            },
            want:    6,
            wantErr: false,
        },
        {
            name:    "case2",
            args:    args{
                arr: []int{1,-2, 3, 10, -4,7, 2,-5},
            },
            want:    18,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Calc(tt.args.arr)
            if (err != nil) != tt.wantErr {
                t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Calc() got = %v, want %v", got, tt.want)
            }
        })
    }
}