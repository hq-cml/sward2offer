package _41_get_mid_num_from_stream

import "testing"

func TestFindMidNum1(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name    string
        args    args
        want    float64
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                arr: []int{4,1,3,2,6,5,7,8},
            },
            want:    4.5,
            wantErr: false,
        },
        {
            name:    "case2",
            args:    args{
                arr: []int{4,1,3,2,6,5,7,8,9},
            },
            want:    5,
            wantErr: false,
        },
        {
            name:    "case3",
            args:    args{
                arr: []int{4,1,3,2,6,5,7,8,1},
            },
            want:    4,
            wantErr: false,
        },
        {
            name:    "case4",
            args:    args{
                arr: []int{4,1,3,1,6,5,7,8},
            },
            want:    4.5,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FindMidNum(tt.args.arr)
            if (err != nil) != tt.wantErr {
                t.Errorf("FindMidNum() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("FindMidNum() got = %v, want %v", got, tt.want)
            }
        })
    }
}