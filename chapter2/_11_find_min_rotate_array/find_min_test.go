package _11_find_min_rotate_array

import "testing"

func TestFindMin(t *testing.T) {
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
               arr: []int{3, 4, 5, 1, 2},
           },
           want:    1,
           wantErr: false,
        },
        {
           name:    "case2",
           args:    args{
               arr: []int{2, 3, 4, 5, 1},
           },
           want:    1,
           wantErr: false,
        },
        {
           name:    "case3",
           args:    args{
               arr: []int{1, 2, 3, 4, 5},
           },
           want:    1,
           wantErr: false,
        },
        {
            name:    "case4",
            args:    args{
                arr: []int{1, 1, 1, 0, 1},
            },
            want:    0,
            wantErr: false,
        },
        {
            name:    "case5",
            args:    args{
                arr: []int{1, 1, 0, 0, 1},
            },
            want:    0,
            wantErr: false,
        },
        {
            name:    "case6",
            args:    args{
                arr: []int{1, 0, 0, 0, 0},
            },
            want:    0,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FindMin(tt.args.arr)
            if (err != nil) != tt.wantErr {
                t.Errorf("FindMin() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("FindMin() got = %v, want %v", got, tt.want)
            }
        })
    }
}