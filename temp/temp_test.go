package temp

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	type args struct {
		arr [][]int
		src int
		dst int
		i   int
		j   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: [][]int{
					{0, 2, 0, 0, 0},
					{0, 2, 2, 2, 0},
					{0, 1, 2, 2, 0},
					{0, 2, 0, 0, 0},
				},
				src: 2,
				dst: 3,
				i:   2,
				j:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Replace(tt.args.arr, tt.args.src, tt.args.dst, tt.args.i, tt.args.j)
			for i := 0; i < len(tt.args.arr); i++ {
				//for j:=0;j <len(tt.args.arr[i]); j++ {
				//	fmt.Print(tt.args.arr[i][j], " ")
				//}
				fmt.Println(tt.args.arr[i])
			}
		})
	}
}
