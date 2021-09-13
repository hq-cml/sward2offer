package replace_double_arr

import (
    "fmt"
    "testing"
)

func TestReplace(t *testing.T) {
    type args struct {
        arr []int
        row int
        col int
        src int
        dst int
        x   int
        y   int
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                arr: []int{
                    0, 2, 0, 0, 0,
                    0, 2, 2, 2, 0,
                    0, 1, 2, 2, 0,
                    0, 2, 0, 0, 0,
                },
                row: 4,
                col: 5,
                src: 2,
                dst: 3,
                x:   2,
                y:   2,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := Replace(tt.args.arr, tt.args.row, tt.args.col, tt.args.src, tt.args.dst, tt.args.x, tt.args.y); (err != nil) != tt.wantErr {
                t.Errorf("Replace() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
        for x:=0; x<tt.args.row; x++ {
            for y:=0; y<tt.args.col; y++ {
                fmt.Print(tt.args.arr[x*tt.args.col + y], " ")
            }
            fmt.Println()
        }

    }
}