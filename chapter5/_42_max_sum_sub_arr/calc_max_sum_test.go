package _42_max_sum_sub_arr

import (
	"testing"
)

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
			name: "case1",
			args: args{
				arr: []int{1, 2, -4, 3, 1, 2, -4},
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				arr: []int{1, -2, 3, 10, -4, 7, 2, -5},
			},
			want:    18,
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				arr: []int{1, -2, 3, 5, -3, 2},
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "case4",
			args: args{
				arr: []int{1, -2, 3, 5, -1, 2},
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "case5",
			args: args{
				arr: []int{-9, -2, -3, -5, -3},
			},
			want:    -2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcMax(tt.args.arr)
			if got != tt.want {
				t.Errorf("Calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
