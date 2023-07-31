package _015_sum_of_3_num

import "testing"

func TestFind3Num(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: []int{-1,0,1,2,-1,-4},
				sum: 0,
			},
		},
		{
			name: "case2",
			args: args{
				arr: []int{0,1,1},
				sum: 0,
			},
		},
		{
			name: "case3",
			args: args{
				arr: []int{0,0,0},
				sum: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Find3Num(tt.args.arr, tt.args.sum)
		})
	}
}